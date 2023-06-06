package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSmt(t *testing.T) {
	stdOut := os.Stdout

	var expected = "epsilon"

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go printSmt(expected, &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)

	output := string(result)

	os.Stdout = stdOut

	fmt.Print(result)

	if !strings.Contains(output, expected) {
		t.Errorf("Expected %s, but not there", expected)
	}

}
