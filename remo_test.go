package remo

import "testing"

func TestNew(t *testing.T) {
	remo := New()
	if remo == nil {
		t.Fatal("remo.New() should be returned pointer of Remo{}")
	}
}
