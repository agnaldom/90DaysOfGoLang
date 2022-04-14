package sayhello

import "testing"

func TestSayHello(t *testing.T) {
	greeting := SayHello()
	if greeting != "Welcome to the world\n" {
		t.Error("TEST FAILED")
	}
}
