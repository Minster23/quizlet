package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// HAPUS KATA "main." -> Panggil langsung fungsinya
	helloHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code salah: dapat %v, ekspektasi %v", status, http.StatusOK)
	}

	ekspektasi := "Hello, world! v5"
	if rr.Body.String() != ekspektasi {
		t.Errorf("Isi body salah: dapat %q, ekspektasi %q", rr.Body.String(), ekspektasi)
	}
}
