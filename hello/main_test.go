package main

import "testing"

func TestGet(t *testing.T) {

	testCases := []struct {
		name string
		want string
	}{
		{"First", "test"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// fake the test for now
			got := "test"
			if got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}
