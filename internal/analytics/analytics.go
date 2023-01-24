package analytics

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"
	"time"

	inputsanitizer "github.com/di-graph/input-sanitizer"
	"github.com/hashicorp/go-uuid"
)

type Analytics interface {
	SetClient(clientFunc func() *http.Client)
	SetCmdArgs(args []string)
	SetVersion(version string)
	GetRequest() (*http.Request, error)
	Send() (*http.Response, error)
	SetTraceId()
	GetTraceId() string
	AddError(err error)
	SetAPIKey(apiKey string)
}

type AnalyticsWrapper struct {
	clientFunc func() *http.Client
	headerFunc func() http.Header

	version   string
	created   time.Time
	args      []string
	traceId   string
	lastError error
	apiKey    string
}

type analyticsData struct {
	Command       string        `json:"command"`
	Subcommand    string        `json:"subcommand"`
	Args          []string      `json:"args"`
	OsPlatform    string        `json:"os_platform"`
	OsArch        string        `json:"os_arch"`
	Os            string        `json:"os"`
	OsRelease     string        `json:"os_release"`
	OsId          string        `json:"os_id"`
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

const ANALYTICS_URL = "https://app.getdigraph.com/api/validate/analytics"

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

func (a *AnalyticsWrapper) SetVersion(version string) {
	a.version = version
}

func (a *AnalyticsWrapper) SetAPIKey(apiKey string) {
	a.apiKey = apiKey
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
	if len(a.args) > 0 {
		output.Command = a.args[0]
	}

	var isCIMode = false
	var commandArgs []string
	if len(a.args) > 1 {
		if !strings.HasPrefix(a.args[1], "--") {
			output.Subcommand = a.args[1]
		}
	}
	for _, arg := range a.args {
		if strings.HasPrefix(arg, "--") {
			if !(strings.Contains(arg, "api-key=") || strings.Contains(arg, "terraform-api-key=")) {
				commandArgs = append(commandArgs, arg)
			}
			if strings.HasPrefix(arg, "--issue-number") || strings.HasPrefix(arg, "--commit-sha") {
				isCIMode = true
			}
		}
	}
	output.Args = commandArgs
	output.Ci = isCIMode

	hsha2 := sha256.Sum256([]byte(a.apiKey))

	output.OsId = fmt.Sprintf("%v", hsha2)

	output.OsPlatform = runtime.GOOS
	output.OsArch = runtime.GOARCH
	output.Version = a.version

	if a.lastError != nil {
		output.ErrorMetadata = ErrorMetadata{ErrorMessage: a.lastError.Error()}
	}

	output.DurationMs = int64(time.Since(a.created).Milliseconds())
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

	request.Header.Set("X-Digraph-Secret-Key", a.apiKey)
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
