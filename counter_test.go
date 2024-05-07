package main

import (
	"testing"
)

func TestCounter_Take(t *testing.T) {
	c := NewCounter(10)

	err := c.Take()
	if err != nil {
		t.Fatalf("Could not take entitlement counter... %v", err)
	}

	if c.Value() != 9 {
		t.Fatalf("counter did not reduce. Expected 9, got %d", c.Value())
	}
}

func TestCounter_TakeError(t *testing.T) {
	c := NewCounter(11)

	err := c.Take()
	if err != nil {
		t.Fatalf("Could not take entitlement counter... %v", err)
	}

	if c.Value() != 9 {
		t.Fatalf("counter did not reduce. Expected 9, got %d", c.Value())
	}
}
