package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("fast one wins", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got, err := Racer(slowServer.URL, fastServer.URL)
		if err != nil {
			t.Fatalf("didn't expect an error but got one:  %v", err)
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
	t.Run("racer fails after set timeout", func(t *testing.T) {
		slowServer := makeDelayedServer(11 * time.Millisecond)
		fastServer := makeDelayedServer(15 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		_, err := ConfigurableRacer(slowServer.URL, fastServer.URL, 10*time.Millisecond)
		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
