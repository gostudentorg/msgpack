package msgpack

import (
	"testing"

	"github.com/matryer/is"
)

func TestConversion_ToNumbers(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		wantInt     int
		wantInt64   int64
		wantUInt64  uint64
		wantFloat64 float64
		wantString  string
	}{
		{
			name:        "nil",
			input:       nil,
			wantInt:     0,
			wantInt64:   0,
			wantUInt64:  0,
			wantFloat64: 0,
			wantString:  "",
		},
		{
			name:        "int",
			input:       3,
			wantInt:     3,
			wantInt64:   3,
			wantUInt64:  3,
			wantFloat64: 3,
			wantString:  "3",
		},
		{
			name:        "int8",
			input:       int8(3),
			wantInt:     3,
			wantInt64:   3,
			wantUInt64:  3,
			wantFloat64: 3,
			wantString:  "3",
		},
		{
			name:        "int16",
			input:       int16(326),
			wantInt:     326,
			wantInt64:   326,
			wantUInt64:  326,
			wantFloat64: 326,
			wantString:  "326",
		},
		{
			name:        "int32",
			input:       int32(2147483647),
			wantInt:     2147483647,
			wantInt64:   2147483647,
			wantUInt64:  2147483647,
			wantFloat64: 2147483647,
			wantString:  "2147483647",
		},
		{
			name:        "int64",
			input:       int64(9223372036854775807),
			wantInt:     9223372036854775807,
			wantInt64:   9223372036854775807,
			wantUInt64:  9223372036854775807,
			wantFloat64: 9223372036854775807,
			wantString:  "9223372036854775807",
		},
		{
			name:        "negative int",
			input:       -3,
			wantInt:     -3,
			wantInt64:   -3,
			wantFloat64: -3,
			wantString:  "-3",
		},
		{
			name:        "negative int8",
			input:       int8(-3),
			wantInt:     -3,
			wantInt64:   -3,
			wantFloat64: -3,
			wantString:  "-3",
		},
		{
			name:        "negative int16",
			input:       int16(-326),
			wantInt:     -326,
			wantInt64:   -326,
			wantFloat64: -326,
			wantString:  "-326",
		},
		{
			name:        "negative int32",
			input:       int32(-2147483647),
			wantInt:     -2147483647,
			wantInt64:   -2147483647,
			wantFloat64: -2147483647,
			wantString:  "-2147483647",
		},
		{
			name:        "negative int64",
			input:       int64(-9223372036854775807),
			wantInt:     -9223372036854775807,
			wantInt64:   -9223372036854775807,
			wantFloat64: -9223372036854775807,
			wantString:  "-9223372036854775807",
		},
		{
			name:        "uint8",
			input:       uint8(3),
			wantInt:     3,
			wantInt64:   3,
			wantUInt64:  3,
			wantFloat64: 3,
			wantString:  "3",
		},
		{
			name:        "uint16",
			input:       uint16(326),
			wantInt:     326,
			wantInt64:   326,
			wantUInt64:  326,
			wantFloat64: 326,
			wantString:  "326",
		},
		{
			name:        "uint32",
			input:       uint32(2147483647),
			wantInt:     2147483647,
			wantInt64:   2147483647,
			wantUInt64:  2147483647,
			wantFloat64: 2147483647,
			wantString:  "2147483647",
		},
		{
			name:        "uint64",
			input:       uint64(9223372036854775807),
			wantInt:     9223372036854775807,
			wantInt64:   9223372036854775807,
			wantUInt64:  9223372036854775807,
			wantFloat64: 9223372036854775807,
			wantString:  "9223372036854775807",
		},
		{
			name:        "float32",
			input:       float32(922337.5625),
			wantInt:     922337,
			wantInt64:   922337,
			wantUInt64:  922337,
			wantFloat64: 922337.5625,
			wantString:  "922337.5625",
		},
		{
			name:        "float64",
			input:       9223372013568.531,
			wantInt:     9223372013568,
			wantInt64:   9223372013568,
			wantUInt64:  9223372013568,
			wantFloat64: 9223372013568.531,
			wantString:  "9223372013568.531",
		},
		{
			name:        "negative float32",
			input:       float32(-922337.5625),
			wantInt:     -922337,
			wantInt64:   -922337,
			wantFloat64: -922337.5625,
			wantString:  "-922337.5625",
		},
		{
			name:        "negative float64",
			input:       -9223372013568.531,
			wantInt:     -9223372013568,
			wantInt64:   -9223372013568,
			wantFloat64: -9223372013568.531,
			wantString:  "-9223372013568.531",
		},
		{
			name:        "string",
			input:       "9223372036854775807",
			wantInt:     9223372036854775807,
			wantInt64:   9223372036854775807,
			wantUInt64:  9223372036854775807,
			wantFloat64: 9223372036854775807,
			wantString:  "9223372036854775807",
		},
		{
			name:        "negative string",
			input:       "-9223372036854775807",
			wantInt:     -9223372036854775807,
			wantInt64:   -9223372036854775807,
			wantFloat64: -9223372036854775807,
			wantString:  "-9223372036854775807",
		},
		{
			name:        "string float",
			input:       "9223372036854.532",
			wantInt:     9223372036854,
			wantInt64:   9223372036854,
			wantUInt64:  9223372036854,
			wantFloat64: 9223372036854.532,
			wantString:  "9223372036854.532",
		},
		{
			name:        "negative string float",
			input:       "-9223372036854.532",
			wantInt:     -9223372036854,
			wantInt64:   -9223372036854,
			wantFloat64: -9223372036854.532,
			wantString:  "-9223372036854.532",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asrt := is.New(t)

			if tt.wantInt != 0 || tt.input == nil {
				asrt.Equal(tt.wantInt, ToInt(tt.input))
			}
			if tt.wantInt64 != 0 || tt.input == nil {
				asrt.Equal(tt.wantInt64, ToInt64(tt.input))
			}
			if tt.wantUInt64 != 0 || tt.input == nil {
				asrt.Equal(tt.wantUInt64, ToUInt64(tt.input))
			}
			if tt.wantFloat64 != 0 || tt.input == nil {
				asrt.Equal(tt.wantFloat64, ToFloat64(tt.input))
			}
			if tt.wantString != "" || tt.input == nil {
				asrt.Equal(tt.wantString, ToString(tt.input))
			}
		})
	}
}
