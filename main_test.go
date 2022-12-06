package main

import "testing"

func TestPing(t *testing.T) {
	got := ping()
  goy := pong()
	if got != "pong" {
		t.Errorf("ping() =%v; want pong", got)
	}
  if goy != "ping" {
		t.Errorf("pong() =%v; want ping", goy)
	}

}
