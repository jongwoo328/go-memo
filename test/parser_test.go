package test

import (
	"go-memo/note/util"
	"testing"
)

func compareStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestParser(t *testing.T) {
	testCase := []struct {
		name          string
		command       string
		wantedCommand string
		wantedOptions []string
	}{
		{"a b c d", "a b c d", "a", []string{"b", "c", "d"}},
		{"a -b --c 1", "a -b --c 1", "a", []string{"-b", "--c", "1"}},
		{"a -b '123' -c 'a b c'", "a -b '123' -c 'a b c'", "a", []string{"-b", "123", "-c", "a b c"}},
		{"a", "a", "a", []string{}},
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			gotCommand, gotOptions := util.ParseCommand(tt.command)
			if gotCommand != tt.wantedCommand || !compareStringSlice(gotOptions, tt.wantedOptions) {
				t.Errorf("got %s %v, want %s %v", gotCommand, gotOptions, tt.wantedCommand, tt.wantedOptions)
			}
		})
	}
}
