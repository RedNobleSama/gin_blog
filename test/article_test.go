package test

import (
	"GinBlog/model"
	"GinBlog/routes"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)
var router *gin.Engine

func init() {
	router = routes.InitRouter()
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestHelloWorld(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"hello": "world",
	}
	// Grab our router
	// Perform a GET request with that handler.
	w := performRequest(router, "GET", "/")
	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)
	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	// Grab the value & whether or not it exists
	value, exists := response["hello"]
	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["hello"], value)
}


func TestInsertArticle(t *testing.T) {
	article := model.Article{
		Title:    "go",
		Cid: 1,
		Content: "hello gin",
		Desc: "hello gin",
	}
	marshal, _ := json.Marshal(article)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/article/add", bytes.NewBufferString(string(marshal)))
	req.Header.Add("content-typ", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.NotEqual(t, "{id:-1}", w.Body.String())
}

func TestLogin(t *testing.T) {
	user := model.User{
		Username: "admin",
		Password: "123456",
	}
	marshal, _ := json.Marshal(user)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBufferString(string(marshal)))
	req.Header.Add("content-typ", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.NotEqual(t, "{id:-1}", w.Body.String())
}