package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct{
		s string
		ans int
	}{
		//normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		//Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbb", 1},
		{"abcabcabcd", 4},

		//chinese
		{"这里啥也不是", 6},
		{"一二三二一", 3},
		{"黑化肥挥发会发灰会花飞灰化肥挥发发黑会飞花",8},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSunStr2(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发会发灰会花飞灰化肥挥发发黑会飞花"
	ans := 8

	for i := 0; i < 13; i++ {
		s = s + s
	}

	b.Logf("len(s) = %d", len(s))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSunStr2(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expected %d", actual, s, ans)
		}
	}
}