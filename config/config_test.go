package config

import (
	"os"
	"testing"
)

func TestParseArgsValid(t *testing.T) {
	originalArgs := os.Args
	t.Cleanup(func() {
		os.Args = originalArgs
	})

	os.Args = []string{"pomodoro", "25", "5", "4"}

	got, err := ParseArgs()

	if err != nil {
		t.Fatalf("ParseArgs() returned an unexpected error %v", err)
	}

	want := Config{WorkTime: 25, RelaxTime: 5, Iterations: 4}

	if got != want {
		t.Errorf("ParseArgs() = %+v, want %+v", got, want)
	}
}

func TestParseArgsRejectsZero(t *testing.T) {
	originalArgs := os.Args
	t.Cleanup(func() {
		os.Args = originalArgs
	})

	os.Args = []string{"pomodoro", "0", "0", "0"}

	_, err := ParseArgs()

	if err == nil {
		t.Fatal("ParseArgs() expected an error for zero, got nil")
	}
}
