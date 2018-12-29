package main

import "testing"

type inAndOut struct {
	limit    int
	expected int
}

var testedValues = []inAndOut{
	{10, 23},
	{1000, 233168},
	{10000, 23331668},
}

func TestSumOf3And5MultsGivenLimit(t *testing.T) {
	for _, test := range testedValues {
		sum := SumOf3And5MultsGivenLimit(test.limit)
		if sum != test.expected {
			t.Errorf("SumOf3And5MultsGivenLimit(%d) = %d, WANT: %d", test.limit, sum, test.expected)
		}
	}
}

func BenchmarkSumOf3And5MultsGivenLimit(b *testing.B) {
	in := 1000000000
	for i := 0; i < b.N; i++ {
		SumOf3And5MultsGivenLimit(in)
	}
}
