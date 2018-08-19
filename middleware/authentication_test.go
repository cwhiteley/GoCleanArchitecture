package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestThatAuthenticationMiddlewareShouldPassControlDownTheChainWhenAuthorizationHeaderIsPresent(t *testing.T) {
	wasInvoked := false

	r, err := http.NewRequest("GET", "/v1/foo", nil)
	require.NoError(t, err, "failed to create a request")
	r.Header.Set("Authorization", "Bearer foobar")
	w := httptest.NewRecorder()

	handlerWithAuthentication := Authenticate(func(w http.ResponseWriter, r *http.Request) {
		wasInvoked = true
		w.WriteHeader(http.StatusOK)
	})

	handlerWithAuthentication(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, wasInvoked)
}

func TestThatAuthenticationMiddlewareShouldReturnA401WhenAuthorizationHeaderIsPresent(t *testing.T) {
	wasInvoked := false

	r, err := http.NewRequest("GET", "/v1/foo", nil)
	require.NoError(t, err, "failed to create a request")
	w := httptest.NewRecorder()

	handlerWithAuthentication := Authenticate(func(w http.ResponseWriter, r *http.Request) {
		wasInvoked = true
		w.WriteHeader(http.StatusOK)
	})

	handlerWithAuthentication(w, r)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.False(t, wasInvoked)
}
