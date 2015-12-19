package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_One(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://localhost:8080/posts/hello", nil)

	response := httptest.NewRecorder()
	p := httprouter.Params{httprouter.Param{"id", "hello"}}
	One(response, r, p)

	exp := "Here is the id you entered hello"
	act := response.Body.String()
	if exp != act {
		t.Fatalf("Expected %s got %s", exp, act)
	}
}
