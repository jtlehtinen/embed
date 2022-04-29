package main

import (
	"bytes"
	"log"
	"os"
	"reflect"
	"testing"
)

var inFilename = "test_in.txt"

func TestMain(m *testing.M) {
	err := os.WriteFile(inFilename, []byte("hello world"), 0644)
	if err != nil {
		log.Fatal(err)
	}

	exitCode := m.Run()

	os.Remove(inFilename)

	os.Exit(exitCode)
}

func TestRun(t *testing.T) {
	args := []string{"test", "-var", "file", "test_in.txt"}
	stdout := bytes.Buffer{}

	err := run(args, &stdout, os.Stderr)
	if err != nil {
		t.Errorf("got unexpected error %q", err)
	}

	want := []byte(
		`const uint8_t file[] = {
  0x68,0x65,0x6c,0x6c,0x6f,0x20,0x77,0x6f,
  0x72,0x6c,0x64,
};
`)

	got := stdout.Bytes()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\n\tgot:  %q\n\twant: %q\n", string(got), string(want))
	}
}
