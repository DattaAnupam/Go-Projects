package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func Test_updateMessage(t *testing.T) {
	msg2 := "update"

	wg.Add(1)
	go updateMessage(msg2, &wg)
	wg.Wait()

	if msg != msg2 {
		t.Errorf("expected: %s, but it wasn't there", msg2)
	}
}

func Test_printMessage(t *testing.T) {
	// store original standard output
	stdOut := os.Stdout

	read, write, _ := os.Pipe()
	os.Stdout = write

	msg = "update"
	printMessage()

	_ = write.Close()

	result, _ := io.ReadAll(read)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, msg) {
		t.Errorf("Expected: %s, but it is not in the output", msg)
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Error("Expected: Hello, universe!, but it is not in the output")
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Error("Expected: Hello, cosmos!, but it is not in the output")
	}

	if !strings.Contains(output, "Hello, world!") {
		t.Error("Expected: Hello, world!, but it is not in the output")
	}
}
