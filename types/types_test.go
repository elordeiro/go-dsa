package types

import (
	"testing"
	"time"
)

func TestIsZero(t *testing.T) {
	type args struct {
		val any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"nil", args{nil}, true},
		{"pointer", args{&args{}}, false},
		{"int", args{int(0)}, true},
		{"int", args{int(1)}, false},
		{"int8", args{int8(0)}, true},
		{"int8", args{int8(1)}, false},
		{"int16", args{int16(0)}, true},
		{"int16", args{int16(1)}, false},
		{"int32", args{int32(0)}, true},
		{"int32", args{int32(1)}, false},
		{"int64", args{int64(0)}, true},
		{"int64", args{int64(1)}, false},
		{"uint", args{uint(0)}, true},
		{"uint", args{uint(1)}, false},
		{"uint8", args{uint8(0)}, true},
		{"uint8", args{uint8(1)}, false},
		{"uint16", args{uint16(0)}, true},
		{"uint16", args{uint16(1)}, false},
		{"uint32", args{uint32(0)}, true},
		{"uint32", args{uint32(1)}, false},
		{"uint64", args{uint64(0)}, true},
		{"uint64", args{uint64(1)}, false},
		{"float32", args{float32(0)}, true},
		{"float32", args{float32(1)}, false},
		{"float64", args{float64(0)}, true},
		{"float64", args{float64(1)}, false},
		{"string", args{""}, true},
		{"string", args{"a"}, false},
		{"bool", args{false}, true},
		{"bool", args{true}, false},
		{"[]any", args{[]any{}}, true},
		{"[]any1", args{[]any{1}}, false},
		{"[]int", args{[]int{}}, true},
		{"[]int1", args{[]int{1}}, false},
		{"time.Time", args{time.Time{}}, true},
		{"time.Time", args{time.Now()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFalsy(tt.args.val); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
