package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go updateMessage("new message", &wg)
	wg.Wait()

	if msg != "new message" {
		t.Errorf("Expected 'new message' got %s", msg)
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "alpha"
	printMessage()

	_ = w.Close()
	data, _ := io.ReadAll(r)

	if !strings.Contains(string(data), "alpha") {
		t.Errorf("expected alpha instead got %s", string(data))
	}

	os.Stdout = stdOut
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()
	data, _ := io.ReadAll(r)

	if !strings.Contains(string(data), "Hello universe") {
		t.Errorf("expected 'Hello universe' instead got %s", string(data))
	}

	if !strings.Contains(string(data), "Hello cosmos") {
		t.Errorf("expected 'Hello cosmos' instead got %s", string(data))
	}

	if !strings.Contains(string(data), "Hello world") {
		t.Errorf("expected 'Hello world' instead got %s", string(data))
	}

	os.Stdout = stdOut
}
