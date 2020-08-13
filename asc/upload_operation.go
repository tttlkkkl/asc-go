package asc

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"sync"
)

// UploadOperations is a slice of UploadOperation instances
type UploadOperations []UploadOperation

// UploadOperation defines model for UploadOperation.
//
// https://developer.apple.com/documentation/appstoreconnectapi/uploadoperation
// https://developer.apple.com/documentation/appstoreconnectapi/uploading_assets_to_app_store_connect
type UploadOperation struct {
	Length         *int                    `json:"length,omitempty"`
	Method         *string                 `json:"method,omitempty"`
	Offset         *int                    `json:"offset,omitempty"`
	RequestHeaders []UploadOperationHeader `json:"requestHeaders,omitempty"`
	URL            *string                 `json:"url,omitempty"`
}

// UploadOperationHeader defines model for UploadOperationHeader.
//
// https://developer.apple.com/documentation/appstoreconnectapi/uploadoperationheader
type UploadOperationHeader struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// UploadOperationError pairs a failed operation and its associated error so it
// can be retried later.
type UploadOperationError struct {
	Operation UploadOperation
	Err       error
}

func (e UploadOperationError) Error() string {
	return e.Err.Error()
}

// chunk returns the bytes in the file from the given offset and with the given length
func (op *UploadOperation) chunk(f *os.File) (io.Reader, error) {
	if op.Offset == nil || op.Length == nil {
		return nil, errors.New("could not establish bounds of upload operation")
	}
	_, err := f.Seek(int64(*op.Offset), 0)
	if err != nil {
		return nil, err
	}
	data := make([]byte, *op.Length)
	_, err = f.Read(data)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(data), nil
}

// request creates a new http.request instance from the given UploadOperation and buffer
func (op *UploadOperation) request(data io.Reader) (*http.Request, error) {
	if op.Method == nil || op.URL == nil {
		return nil, errors.New("could not establish destination of upload operation")
	}
	req, err := http.NewRequest(*op.Method, *op.URL, data)
	if err != nil {
		return nil, err
	}
	if op.RequestHeaders != nil {
		for _, h := range op.RequestHeaders {
			if h.Name == nil || h.Value == nil {
				continue
			}
			req.Header.Add(*h.Name, *h.Value)
		}
	}
	return req, nil
}

// Upload takes a file path and concurrently uploads each part of the file to App Store Connect
func (ops UploadOperations) Upload(ctx context.Context, file *os.File, client *Client) error {
	var wg sync.WaitGroup
	errs := make(chan UploadOperationError)

	for i, operation := range ops {
		chunk, err := operation.chunk(file)
		if err != nil {
			errs <- UploadOperationError{
				Operation: operation,
				Err:       err,
			}
			continue
		}
		wg.Add(1)
		go client.sendChunk(ctx, ops[i], chunk, errs, &wg)
	}

	go func() {
		wg.Wait()
		close(errs)
	}()

	for err := range errs {
		return err
	}

	return nil
}

func (c *Client) sendChunk(ctx context.Context, op UploadOperation, chunk io.Reader, errs chan<- UploadOperationError, wg *sync.WaitGroup) {
	defer wg.Done()
	req, err := op.request(chunk)
	if err != nil {
		errs <- UploadOperationError{
			Operation: op,
			Err:       err,
		}
		return
	}
	req = req.WithContext(ctx)
	_, err = c.do(ctx, req, nil)
	if err != nil {
		errs <- UploadOperationError{
			Operation: op,
			Err:       err,
		}
	}

}
