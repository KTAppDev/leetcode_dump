package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	expect := "fl"
	got := longestCommonPrefix([]string{"flower", "flow", "flight"})
	if got != expect {
		t.Errorf("Got %s; Want %s", got, expect)
	}
}
