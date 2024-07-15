// checker/checker_test.go
package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func runChecker(t *testing.T, args string, instructions string) string {
	cmd := exec.Command("go", "run", "../checker/checker.go", args)
	cmd.Stdin = strings.NewReader(instructions)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		t.Errorf("Error running command: %v", err)
	}

	return out.String()
}

func TestSortedInputWithNoInstructions(t *testing.T) {
	output := runChecker(t, "1 2 3 4", "")
	expected := "OK\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestReverseSortedInputWithProperInstructions(t *testing.T) {
	output := runChecker(t, "3 2 1 0", "rra\npb\nsa\nrra\npa\n")
	expected := "OK\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestUnsortedInputWithNoInstructions(t *testing.T) {
	output := runChecker(t, "3 2 1 4", "")
	expected := "KO\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestInvalidInstruction(t *testing.T) {
	output := runChecker(t, "3 2 1 4", "invalid\n")
	expected := "Error\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestDuplicateNumberInInput(t *testing.T) {
	output := runChecker(t, "1 2 2 3", "")
	expected := "Error\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestNonIntegerInput(t *testing.T) {
	output := runChecker(t, "1 2 three 4", "")
	expected := "Error\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestMixedInstructions(t *testing.T) {
	output := runChecker(t, "3 2 1 0", "sa\nrra\npb\n")
	expected := "KO\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestCheckerWithCorrectInstructions(t *testing.T) {
	output := runChecker(t, "3 2 1 0", "rra\npb\nsa\nrra\npa\n")
	expected := "OK\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestCheckerWithPartialInstructions(t *testing.T) {
	output := runChecker(t, "3 2 1 0", "sa\n")
	expected := "KO\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestCheckerWithExtraInstructions(t *testing.T) {
	output := runChecker(t, "3 2 1 0", "sa\nrra\npb\npa\nra\n")
	expected := "KO\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}
