package twosum

import "testing"

func TestTwoSum(t *testing.T) {
	for _, tc := range testCases {
		result := twoSum(tc.nums, tc.target)
		if tc.result == nil && result != nil{
			t.Fatalf("twoSum(%v, %d) = %v, want nil", tc.nums, tc.target, result)
		} else if tc.result != nil && result == nil {
			t.Fatalf("twoSum(%v, %d) = nil, want %v", tc.nums, tc.target, tc.result)
		} else if tc.result != nil && result != nil {
			if tc.result[0] != result[0] || tc.result[1] != result[1] {
				t.Fatalf("twoSum(%v, %d) = %v, want %v", tc.nums, tc.target, result, tc.result)
			}
		}
	}
}

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			twoSum(tc.nums, tc.target)
		}
	}
}