package utils

import (
	"fmt"
	"testing"
)

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

func ExampleStrToFloat64() {
	out1, _ := StrToFloat64("3.14")
	fmt.Println(out1)
	out2, _ := StrToFloat64("-10000")
	fmt.Println(out2)
	out3, _ := StrToFloat64("0.00000001")
	fmt.Println(out3)

	// Output:
	// 3.14
	// -10000
	// 1e-08
}

func ExampleStrToInt64() {
	out1, _ := StrToInt64("3")
	fmt.Println(out1)
	out2, _ := StrToInt64("-10000")
	fmt.Println(out2)
	out3, _ := StrToInt64("10000000")
	fmt.Println(out3)

	// Output:
	// 3
	// -10000
	// 10000000
}

func ExampleFloat64ToStr() {
	out1 := Float64ToStr(3.14)
	fmt.Println(out1)
	out2 := Float64ToStr(-10000)
	fmt.Println(out2)
	out3 := Float64ToStr(0.0000001)
	fmt.Println(out3)

	// Output:
	// 3.14
	// -10000
	// 0.0000001
}

func ExampleInt64ToStr() {
	out1 := Int64ToStr(314)
	fmt.Println(out1)
	out2 := Int64ToStr(-10000)
	fmt.Println(out2)
	out3 := Int64ToStr(10000000)
	fmt.Println(out3)

	// Output:
	// 314
	// -10000
	// 10000000
}
