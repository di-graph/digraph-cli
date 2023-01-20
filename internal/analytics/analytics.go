package analytics

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"time"

	inputsanitizer "github.com/di-graph/input-sanitizer"
	"github.com/hashicorp/go-uuid"
)

type Analytics interface {
	SetClient(clientFunc func() *http.Client)
	SetCmdArgs(args []string)
	SetOrg(org string)
	SetVersion(version string)
	GetRequest() (*http.Request, error)
	Send() (*http.Response, error)
	SetTraceId()
	GetTraceId() string
	AddError(err error)
}

type AnalyticsWrapper struct {
	clientFunc func() *http.Client
	headerFunc func() http.Header

	org       string
	version   string
	created   time.Time
	args      []string
	traceId   string
	lastError error
}

type analyticsData struct {
	Command       string        `json:"command"`
	Args          []string      `json:"args"`
	OsPlatform    string        `json:"os_platform"`
	OsArch        string        `json:"os_arch"`
	Os            string        `json:"os"`
	OsRelease     string        `json:"os_release"`
	Id            string        `json:"id"`
	Version       string        `json:"version"`
	DurationMs    int64         `json:"duration_ms"`
	Ci            bool          `json:"ci"`
	ErrorMetadata ErrorMetadata `json:"error_metadata"`
}

type ErrorMetadata struct {
	ErrorMessage string `json:"error_message"`
}

type dataOutput struct {
	Data analyticsData `json:"data"`
}

const ANALYTICS_URL = ""

func New() Analytics {
	a := &AnalyticsWrapper{}
	a.headerFunc = func() http.Header { return http.Header{} }
	a.created = time.Now()
	a.clientFunc = func() *http.Client { return &http.Client{} }
	return a
}

func (a *AnalyticsWrapper) SetCmdArgs(args []string) {
	a.args = args
}

func (a *AnalyticsWrapper) SetOrg(org string) {
	a.org = org
}

func (a *AnalyticsWrapper) SetVersion(version string) {
	a.version = version
}

func (a *AnalyticsWrapper) SetClient(clientFunc func() *http.Client) {
	a.clientFunc = clientFunc
}

func (a *AnalyticsWrapper) GetTraceId() string {
	return a.traceId
}

func (a *AnalyticsWrapper) SetTraceId() {
	// deepcode ignore InsecureHash: It is just being used to generate an id, without any security concerns
	shasum := sha1.New()
	uuid, _ := uuid.GenerateUUID()
	io.WriteString(shasum, uuid)
	a.traceId = fmt.Sprintf("%x", shasum.Sum(nil))
}

func (a *AnalyticsWrapper) AddError(err error) {
	a.lastError = err
}

func (a *AnalyticsWrapper) GetOutputData() *analyticsData {
	output := &analyticsData{}

	output.Id = a.GetTraceId()

	output.Args = a.args[1:]
	if len(a.args) > 0 {
		output.Command = a.args[0]
	}

	output.OsPlatform = runtime.GOOS
	output.OsArch = runtime.GOARCH
	output.Version = a.version
	output.DurationMs = int64(time.Since(a.created).Milliseconds())

	if a.lastError != nil {
		output.ErrorMetadata = ErrorMetadata{ErrorMessage: a.lastError.Error()}
	}

	return output
}

func (a *AnalyticsWrapper) GetRequest() (*http.Request, error) {
	output := a.GetOutputData()

	outputJson, err := json.Marshal(dataOutput{Data: *output})
	if err != nil {
		return nil, err
	}

	outputJson, err = inputsanitizer.SanitizeValuesByKey(outputJson)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(outputJson)
	request, err := http.NewRequest(http.MethodPost, ANALYTICS_URL, body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	return request, err
}

func (a *AnalyticsWrapper) Send() (*http.Response, error) {
	request, err := a.GetRequest()
	if err != nil {
		return nil, err
	}

	client := a.clientFunc()
	response, err := client.Do(request)

	return response, err
}
