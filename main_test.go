package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestGetCounters(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/counters", nil)
	router.ServeHTTP(w, req)

	var c []counter

	err := json.Unmarshal(w.Body.Bytes(), &c)
	require.NoError(t, err)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, len(c), len(counters))
}

func TestGetCounterByID(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/counter/d09b11a1-3ef8-47f6-a4de-620e7cabdc1a", nil)
	router.ServeHTTP(w, req)

	var c counter

	err := json.Unmarshal(w.Body.Bytes(), &c)
	require.NoError(t, err)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 100, c.Value)
}

func TestCreateCounter(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/counter", nil)
	router.ServeHTTP(w, req)

	var c counter

	err := json.Unmarshal(w.Body.Bytes(), &c)
	require.NoError(t, err)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, 0, c.Value)
}

func TestIncrementCounter(t *testing.T) {
	router := setupRouter()

	w1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/counter/5ca44aab-ee12-4911-925c-329175c0d1a0", nil)
	router.ServeHTTP(w1, req1)

	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("POST", "/counter/5ca44aab-ee12-4911-925c-329175c0d1a0", nil)
	router.ServeHTTP(w2, req2)

	var c1 counter
	var c2 counter

	err1 := json.Unmarshal(w1.Body.Bytes(), &c1)
	require.NoError(t, err1)

	err2 := json.Unmarshal(w2.Body.Bytes(), &c2)
	require.NoError(t, err2)

	assert.Equal(t, 200, w1.Code)
	assert.Equal(t, 51, c1.Value)

	assert.Equal(t, 200, w2.Code)
	assert.Equal(t, 52, c2.Value)
}

func TestDeleteCounter(t *testing.T) {
	router := setupRouter()

	w1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("DELETE", "/counter/5ca44aab-ee12-4911-925c-329175c0d1a0", nil)
	router.ServeHTTP(w1, req1)

	var c counter

	err := json.Unmarshal(w1.Body.Bytes(), &c)
	require.NoError(t, err)

	assert.Equal(t, 200, w1.Code)

	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/counter/5ca44aab-ee12-4911-925c-329175c0d1a0", nil)
	router.ServeHTTP(w2, req2)

	assert.Equal(t, 404, w2.Code)
}
