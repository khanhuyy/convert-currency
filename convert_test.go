package convert

import (
	"testing"
)

func TestToWords(t *testing.T) {
	tests := []struct {
		name     string
		input    int64
		expected string
	}{
		{
			name:     "Single number 1",
			input:    0,
			expected: "Không đồng",
		},
		{
			name:     "Single number 2",
			input:    5,
			expected: "Năm đồng",
		},
		{
			name:     "test ten",
			input:    10,
			expected: "Mười đồng",
		},
		{
			name:     "test 45",
			input:    45,
			expected: "Bốn mươi năm đồng",
		},
		{
			name:     "test 1000",
			input:    1000,
			expected: "Một nghìn đồng",
		},
		{
			name:     "test 1012",
			input:    1012,
			expected: "Một nghìn không trăm mười hai đồng",
		},
		{
			name:     "test 45021",
			input:    45021,
			expected: "Bốn mươi năm nghìn không trăm hai mươi mốt đồng",
		},
		{
			name:     "test 45000",
			input:    45000,
			expected: "Bốn mươi năm nghìn đồng",
		},
		{
			name:     "test 400000",
			input:    400000,
			expected: "Bốn trăm nghìn đồng",
		},
		{
			name:     "test 450000",
			input:    450000,
			expected: "Bốn trăm năm mươi nghìn đồng",
		},
		{
			name:     "test 4500000",
			input:    4500000,
			expected: "Bốn triệu năm trăm nghìn đồng",
		},
		{
			name:     "test 7000000",
			input:    7000000,
			expected: "Bảy triệu đồng",
		},
		{
			name:     "test 21021021",
			input:    21021021,
			expected: "Hai mươi mốt triệu không trăm hai mươi mốt nghìn không trăm hai mươi mốt đồng",
		},
		{
			name:     "test 3000000005",
			input:    3000000005,
			expected: "Ba tỷ lẻ năm đồng",
		},
		{
			name:     "test 7000000000",
			input:    7000000000,
			expected: "Bảy tỷ đồng",
		},
		{
			name:     "test 007",
			input:    007,
			expected: "Bảy đồng",
		},
		// todo
		//{
		//	name:     "Min integer",
		//	input:    9223372036854775807,
		//	expected: "Chín tỷ hai trăm hai mươi ba triệu ba trăm bảy mươi hai nghìn không trăm ba mươi sáu tỷ tám trăm năm mươi tư triệu bảy trăm bảy mươi lăm nghìn tám trăm lẻ bảy đồng",
		//},

	}

	for _, test := range tests {
		println("test %s", test.name)
		println(test.input, test.expected)
		output := ToWords(test.input)
		if test.expected != output {
			t.Error("c2 (which should be 3.5) plus 1.2 does not equal 4.7; value:", output)
		}
	}

}
