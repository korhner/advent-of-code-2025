package main

import "testing"

func TestInitialRotation(t *testing.T) {
	c := &Clock{}
	c.Rotate(10)
	if c.CurrentState != 10 {
		t.Errorf("expected 10, got %d", c.CurrentState)
	}
}

func TestContinuousRotation(t *testing.T) {
	c := &Clock{}
	c.Rotate(10)
	c.Rotate(10)
	if c.CurrentState != 20 {
		t.Errorf("expected 20, got %d", c.CurrentState)
	}
}

func TestRolloverRight(t *testing.T) {
	c := &Clock{}
	c.Rotate(99)
	c.Rotate(1)
	if c.CurrentState != 0 {
		t.Errorf("expected 0, got %d", c.CurrentState)
	}
}

func TestRotateLeft(t *testing.T) {
	c := &Clock{}
	c.Rotate(-1)
	if c.CurrentState != 99 {
		t.Errorf("expected 99, got %d", c.CurrentState)
	}
}

func TestRolloverLeft(t *testing.T) {
	c := &Clock{}
	c.Rotate(-101)
	if c.CurrentState != 99 {
		t.Errorf("expected 99, got %d", c.CurrentState)
	}
}
