package asc

import (
	"context"
	"fmt"
)

// DiagnosticLog defines model for DiagnosticLog.
type DiagnosticLog struct {
	ID    string        `json:"id"`
	Links ResourceLinks `json:"links"`
	Type  string        `json:"type"`
}

// DiagnosticLogsResponse defines model for DiagnosticLogsResponse.
type DiagnosticLogsResponse struct {
	Data  []DiagnosticLog    `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// DiagnosticSignature defines model for DiagnosticSignature.
type DiagnosticSignature struct {
	Attributes *struct {
		DiagnosticType *string  `json:"diagnosticType,omitempty"`
		Signature      *string  `json:"signature,omitempty"`
		Weight         *float32 `json:"weight,omitempty"`
	} `json:"attributes,omitempty"`
	ID    string        `json:"id"`
	Links ResourceLinks `json:"links"`
	Type  string        `json:"type"`
}

// DiagnosticSignaturesResponse defines model for DiagnosticSignaturesResponse.
type DiagnosticSignaturesResponse struct {
	Data     []DiagnosticSignature `json:"data"`
	Included *[]DiagnosticLog      `json:"included,omitempty"`
	Links    PagedDocumentLinks    `json:"links"`
	Meta     *PagingInformation    `json:"meta,omitempty"`
}

// PerfPowerMetric defines model for PerfPowerMetric.
type PerfPowerMetric struct {
	Attributes *struct {
		DeviceType *string `json:"deviceType,omitempty"`
		MetricType *string `json:"metricType,omitempty"`
		Platform   *string `json:"platform,omitempty"`
	} `json:"attributes,omitempty"`
	ID    string        `json:"id"`
	Links ResourceLinks `json:"links"`
	Type  string        `json:"type"`
}

// PerfPowerMetricsResponse defines model for PerfPowerMetricsResponse.
type PerfPowerMetricsResponse struct {
	Data  []PerfPowerMetric  `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// GetPerfPowerMetricsQuery are query options for GetPerfPowerMetrics
type GetPerfPowerMetricsQuery struct {
	FilterDeviceType []string `url:"filter[deviceType],omitempty"`
	FilterMetricType []string `url:"filter[metricType],omitempty"`
	FilterPlatform   []string `url:"filter[platform],omitempty"`
	Cursor           string   `url:"cursor,omitempty"`
}

// ListDiagnosticsSignaturesQuery are query options for ListDiagnosticsSignatures
type ListDiagnosticsSignaturesQuery struct {
	FieldsDiagnosticSignatures []string `url:"fields[diagnosticSignatures],omitempty"`
	FilterDiagnosticType       []string `url:"filter[diagnosticType],omitempty"`
	Limit                      int      `url:"limit,omitempty"`
	Cursor                     string   `url:"cursor,omitempty"`
}

// GetLogsForDiagnosticSignatureQuery are query options for GetLogsForDiagnosticSignature
type GetLogsForDiagnosticSignatureQuery struct {
	Limit  int    `url:"limit,omitempty"`
	Cursor string `url:"cursor,omitempty"`
}

// GetPerfPowerMetricsForApp gets the performance and power metrics data for the most recent versions of an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_power_and_performance_metrics_for_an_app
func (s *ReportingService) GetPerfPowerMetricsForApp(ctx context.Context, id string, params *GetPerfPowerMetricsQuery) (*PerfPowerMetricsResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/perfPowerMetrics", id)
	res := new(PerfPowerMetricsResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetPerfPowerMetricsForBuild gets the performance and power metrics data for a specific build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_power_and_performance_metrics_for_a_build
func (s *ReportingService) GetPerfPowerMetricsForBuild(ctx context.Context, id string, params *GetPerfPowerMetricsQuery) (*PerfPowerMetricsResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s/perfPowerMetrics", id)
	res := new(PerfPowerMetricsResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// ListDiagnosticSignaturesForBuild lists the aggregate backtrace signatures captured for a specific build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_diagnostic_signatures_for_a_build
func (s *ReportingService) ListDiagnosticSignaturesForBuild(ctx context.Context, id string, params *ListDiagnosticsSignaturesQuery) (*DiagnosticSignaturesResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s/diagnosticSignatures", id)
	res := new(DiagnosticSignaturesResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetLogsForDiagnosticSignature gets the anonymized backtrace logs associated with a specific diagnostic signature.
//
// https://developer.apple.com/documentation/appstoreconnectapi/download_logs_for_a_diagnostic_signature
func (s *ReportingService) GetLogsForDiagnosticSignature(ctx context.Context, id string, params *GetLogsForDiagnosticSignatureQuery) (*DiagnosticLogsResponse, *Response, error) {
	url := fmt.Sprintf("diagnosticSignatures/%s/logs", id)
	res := new(DiagnosticLogsResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}
