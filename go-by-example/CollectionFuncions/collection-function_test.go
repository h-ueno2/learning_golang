package main

import "testing"

func TestIndex(t *testing.T) {
	strs := []string{"peach", "apple", "pear", "plum"}

	tests := []struct {
		name string
		t    string
		want int
	}{
		{
			name: "normal",
			t:    "pear",
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Index(strs, tt.t)
			if got != tt.want {
				t.Errorf("Index() = %d, want %d", got, tt.want)
			}
		})
	}
}
