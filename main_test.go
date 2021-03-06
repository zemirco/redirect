package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/zemirco/redirect/router"
)

func TestIndex(t *testing.T) {
	server := httptest.NewServer(router.NewRouter())
	defer server.Close()
	res, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Error("invalid status code")
	}
	if !strings.Contains(string(body), "Hi there!") {
		t.Error("invalid body")
	}
}

func TestLogin(t *testing.T) {
	server := httptest.NewServer(router.NewRouter())
	defer server.Close()
	res, err := http.PostForm(server.URL+"/login", url.Values{})
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Error("invalid status code")
	}
	if !strings.Contains(string(body), "Hi there!") {
		t.Error("invalid body")
	}
}
