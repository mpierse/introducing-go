package math

import (
	"testing"
)

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

func TestMin(t *testing.T) {
	var v float64
	v = Min([]float64{1, 2})
	if v != 1 {
		t.Error("Expected 1, got ", v)
	}
}

func TestMax(t *testing.T) {
	var v float64
	v = Max([]float64{1, 2})
	if v != 2 {
		t.Error("Expected 2, got ", v)
	}
}

//Table driven tests

type testpair struct {
	values []float64
	answer float64
}

var averageTests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
	{[]float64{}, 0},
}

func TestAverage2(t *testing.T) {
	for _, pair := range averageTests {
		v := Average(pair.values)
		if v != pair.answer {
			t.Error("Expected ", pair.answer, ", got ", v)
		}
	}
}

var minTests = []testpair{
	{[]float64{1, 2}, 1},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, -1},
	{[]float64{}, 0},
}

func TestMin2(t *testing.T) {
	for _, pair := range minTests {
		v := Min(pair.values)
		if v != pair.answer {
			t.Error("Expected ", pair.answer, ", got ", v)
		}
	}
}

var maxTests = []testpair{
	{[]float64{1, 2}, 2},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 1},
	{[]float64{}, 0},
}

func TestMax2(t *testing.T) {
	for _, pair := range maxTests {
		v := Max(pair.values)
		if v != pair.answer {
			t.Error("Expected ", pair.answer, ", got ", v)
		}
	}
}
