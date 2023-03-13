package utils

import "testing"

func TestIsSameEmpty(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want bool
	}{
		{
			name: "all items not nullable",
			args: []string{"1", "2", "3"},
			want: false,
		},
		{
			name: "all items nullable",
			args: []string{"", "", ""},
			want: true,
		},
		{
			name: "same item  nullable",
			args: []string{"1", "", "3"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSameEmpty(tt.args); got != tt.want {
				t.Errorf("IsSameEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
