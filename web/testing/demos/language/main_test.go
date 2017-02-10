package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Base(t *testing.T) {

	m := createHandlers()

	ts := httptest.NewServer(m)
	defer ts.Close()

	tests := []struct {
		url      string
		expected string
	}{
		{url: "/spa", expected: "Hola Mundo"},
		{url: "/", expected: "Hello World!"},
		{url: "/prt", expected: "Olá Mundo"},
		{url: "/fre", expected: "Bonjour le monde"},
	}

	for _, test := range tests {

		testURL := fmt.Sprintf(ts.URL + test.url)

		// Call the handler through the server's URL.
		res, err := http.Get(testURL)
		if err != nil {
			t.Fatal(err)
		}

		// Read in the response from the call.
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		// Validate we received all the known customers.
		got := string(b)
		if got != test.expected {
			t.Log("Wanted:", test.expected)
			t.Log("Got   :", got)
			t.Fatal("Mismatch")
		}

	}

}
