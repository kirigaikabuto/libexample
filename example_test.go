package libexample

import "testing"

func TestOutput(t *testing.T) {
	result := Output()
	if result != "Hello world" {
		t.Errorf("Was error")
	}
}
