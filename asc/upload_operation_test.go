package asc

import (
	"bytes"
	"context"
	"crypto/rand"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultipartUpload(t *testing.T) {
	file, err := ioutil.TempFile("", "big_file")
	if err != nil {
		assert.FailNow(t, "temp file creation produced an error", err)
	}
	defer os.Remove(file.Name())

	contents := make([]byte, 64)
	_, err = rand.Read(contents)
	if err != nil {
		assert.FailNow(t, "random bytes produced an error", err)
	}
	_, err = file.Write(contents)
	if err != nil {
		assert.FailNow(t, "writing the temp file produced an error", err)
	}

	client, server := newServer("")
	defer server.Close()

	operations := UploadOperations{
		{
			URL:            String(client.BaseURL.String()),
			Offset:         Int(0),
			Length:         Int(10),
			Method:         String("PATCH"),
			RequestHeaders: []UploadOperationHeader{},
		},
		{
			URL:            String(client.BaseURL.String()),
			Offset:         Int(10),
			Length:         Int(10),
			Method:         String("PATCH"),
			RequestHeaders: []UploadOperationHeader{},
		},
		{
			URL:            String(client.BaseURL.String()),
			Offset:         Int(20),
			Length:         Int(30),
			Method:         String("PATCH"),
			RequestHeaders: []UploadOperationHeader{},
		},
		{
			URL:            String(client.BaseURL.String()),
			Offset:         Int(50),
			Length:         Int(10),
			Method:         String("PATCH"),
			RequestHeaders: []UploadOperationHeader{},
		},
		{
			URL:            String(client.BaseURL.String()),
			Offset:         Int(60),
			Length:         Int(4),
			Method:         String("PATCH"),
			RequestHeaders: []UploadOperationHeader{},
		},
	}

	err = operations.Upload(context.Background(), file, client)
	assert.NoError(t, err)
}

func TestUploadOperationChunk(t *testing.T) {
	file, err := ioutil.TempFile("", "small_file")
	if err != nil {
		assert.FailNow(t, "temp file creation produced an error", err)
	}
	defer os.Remove(file.Name())

	contents := make([]byte, 20)
	_, err = rand.Read(contents)
	if err != nil {
		assert.FailNow(t, "random bytes produced an error", err)
	}
	_, err = file.Write(contents)
	if err != nil {
		assert.FailNow(t, "writing the temp file produced an error", err)
	}

	op := UploadOperation{
		URL:            String("test"),
		Offset:         Int(0),
		Length:         Int(10),
		Method:         String("PATCH"),
		RequestHeaders: []UploadOperationHeader{},
	}

	chunk, err := op.chunk(file)
	assert.NoError(t, err)

	buf := &bytes.Buffer{}
	written, err := io.Copy(buf, chunk)
	assert.NoError(t, err)
	assert.EqualValues(t, written, *op.Length)
}
