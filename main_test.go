package main

import (
	"encoding/json"
	"fmt"
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

	var c []Counter

	err := json.Unmarshal(w.Body.Bytes(), &c)
	require.NoError(t, err)

	assert.Equal(t, 200, w.Code)
}

func TestGetCounterByID(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/counter/d09b11a1-3ef8-47f6-a4de-620e7cabdc1a", nil)
	router.ServeHTTP(w, req)

	var c Counter

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

	var c Counter

	err := json.Unmarshal(w.Body.Bytes(), &c)
	require.NoError(t, err)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, 0, c.Value)
}

func TestIncrementCounter(t *testing.T) {
	router := setupRouter()

	w1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/counter", nil)
	router.ServeHTTP(w1, req1)

	var c1 Counter
	err1 := json.Unmarshal(w1.Body.Bytes(), &c1)
	require.NoError(t, err1)

	url := fmt.Sprintf("/counter/%s", c1.ID)

	w2 := httptest.NewRecorder()
	w3 := httptest.NewRecorder()
	req2, _ := http.NewRequest("POST", url, nil)
	req3, _ := http.NewRequest("POST", url, nil)
	router.ServeHTTP(w2, req2)
	router.ServeHTTP(w3, req3)

	var c2 Counter

	err2 := json.Unmarshal(w3.Body.Bytes(), &c2)
	require.NoError(t, err2)

	assert.Equal(t, 200, w3.Code)
	assert.Equal(t, 2, c2.Value)
}

func TestDeleteCounter(t *testing.T) {
	router := setupRouter()

	w1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/counter", nil)
	router.ServeHTTP(w1, req1)

	var c1 Counter
	err1 := json.Unmarshal(w1.Body.Bytes(), &c1)
	require.NoError(t, err1)

	url := fmt.Sprintf("/counter/%s", c1.ID)

	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("DELETE", url, nil)
	router.ServeHTTP(w2, req2)

	var c2 Counter

	err := json.Unmarshal(w1.Body.Bytes(), &c2)
	require.NoError(t, err)

	assert.Equal(t, 200, w2.Code)

	w3 := httptest.NewRecorder()
	req3, _ := http.NewRequest("GET", "/counter/5ca44aab-ee12-4911-925c-329175c0d1a0", nil)
	router.ServeHTTP(w3, req3)

	assert.Equal(t, 404, w3.Code)
}
