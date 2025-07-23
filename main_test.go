package main

import "testing"

func TestExample(t *testing.T) {
    got := 1 + 1
    want := 2
    if got != want {
        t.Errorf("Expected %d, got %d", want, got)
    }
}
