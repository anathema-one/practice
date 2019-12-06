package leetcode

import "testing"

func TestReverseInt(t *testing.T) {
	input := 123
	output := 321

	result := reverse(input)
	if output != result {
		t.Error("expected", output, "got", result)
	}
}
