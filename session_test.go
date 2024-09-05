package veriff_test

import (
	"bytes"
	"context"
	_ "embed"
	"io"
	"net/http"
	"testing"

	"github.com/brokeyourbike/veriff-api-client-go"
	"github.com/sirupsen/logrus"
	logrustest "github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/decision.json
var decisionMsg []byte

//go:embed testdata/veriff-create-session-success.json
var veriffCreateSessionSuccess []byte

func TestCreateSession(t *testing.T) {
	mockHttpClient := veriff.NewMockHttpClient(t)

	logger, hook := logrustest.NewNullLogger()
	logger.SetLevel(logrus.DebugLevel)

	client := veriff.NewClient("https://a.b", "token", "secret", veriff.WithHTTPClient(mockHttpClient), veriff.WithLogger(logger))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(veriffCreateSessionSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.CreateSession(context.TODO(), veriff.CreateSessionPayload{})
	require.NoError(t, err)
	assert.Equal(t, "success", got.Status)

	require.Equal(t, 2, len(hook.Entries))
	require.Contains(t, hook.Entries[0].Data, "http.request.method")
	require.Contains(t, hook.Entries[0].Data, "http.request.url")
	require.Contains(t, hook.Entries[0].Data, "http.request.body.content")
	require.Contains(t, hook.Entries[1].Data, "http.response.status_code")
	require.Contains(t, hook.Entries[1].Data, "http.response.body.content")
	require.Contains(t, hook.Entries[1].Data, "http.response.headers")
}

func TestCreateSession_RequestErr(t *testing.T) {
	mockHttpClient := veriff.NewMockHttpClient(t)
	client := veriff.NewClient("https://a.b", "token", "secret", veriff.WithHTTPClient(mockHttpClient))

	_, err := client.CreateSession(nil, veriff.CreateSessionPayload{}) //lint:ignore SA1012 testing failure
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to create request")
}

func TestSessionDecision(t *testing.T) {
	mockHttpClient := veriff.NewMockHttpClient(t)
	client := veriff.NewClient("https://a.b", "token", "secret", veriff.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(decisionMsg))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.SessionDecision(context.TODO(), "")
	require.NoError(t, err)
	assert.Equal(t, "success", got.Status)
}

func TestSessionDecision_RequestErr(t *testing.T) {
	mockHttpClient := veriff.NewMockHttpClient(t)
	client := veriff.NewClient("https://a.b", "token", "secret", veriff.WithHTTPClient(mockHttpClient))

	_, err := client.SessionDecision(nil, "") //lint:ignore SA1012 testing failure
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to create request")
}

func TestSessionMedia_RequestErr(t *testing.T) {
	mockHttpClient := veriff.NewMockHttpClient(t)
	client := veriff.NewClient("https://a.b", "token", "secret", veriff.WithHTTPClient(mockHttpClient))

	_, err := client.SessionMedia(nil, "") //lint:ignore SA1012 testing failure
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to create request")
}
