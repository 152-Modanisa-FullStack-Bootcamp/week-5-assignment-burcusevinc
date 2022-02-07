package assignment

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestAddUint32(t *testing.T) {
	tests := []struct {
		param1, param2, want uint32
		boolwant             bool
	}{
		{math.MaxUint32, 1, 0, true},
		{1, 1, 2, false},
		{42, 2701, 2743, false},
		{42, math.MaxUint32, 41, true},
		{4294967290, 5, 4294967295, false},
		{4294967290, 6, 0, true},
		{4294967290, 10, 4, true},
	}
	for _, test := range tests {
		sum, overflow := AddUint32(test.param1, test.param2)
		assert.Equal(t, test.want, sum)
		assert.Equal(t, overflow, test.boolwant)
	}

}

func TestCeilNumber1(t *testing.T) {
	tests := []struct {
		params float64 //given params
		want   float64 //expected value
	}{
		//test cases
		{42.42, 42.50},
		{42, 42},
		{42.01, 42.25},
		{42.24, 42.25},
		{42.25, 42.25},
		{42.26, 42.50},
		{42.55, 42.75},
		{42.75, 42.75},
		{42.76, 43},
		{42.99, 43},
		{43.13, 43.25},
	}
	for _, v := range tests {
		point := CeilNumber(v.params) //My Function
		assert.Equal(t, v.want, point)
	}
}

func TestAlphabetSoup1(t *testing.T) {
	tests := []struct {
		params string
		want   string
	}{
		{"hello", "ehllo"},
		{"", ""},
		{"h", "h"},
		{"ab", "ab"},
		{"ba", "ab"},
		{"bac", "abc"},
		{"cba", "abc"},
	}
	for _, v := range tests {
		result := AlphabetSoup(v.params)
		assert.Equal(t, v.want, result)
	}
}

func TestStringMask(t *testing.T) {
	tests := []struct {
		params string
		star   uint
		want   string
	}{
		{"!mysecret*", 2, "!m********"},
		{"", 4, "*"},
		{"a", 1, "*"},
		{"string", 0, "******"},
		{"string", 3, "str***"},
		{"string", 5, "strin*"},
		{"string", 6, "******"},
		{"string", 7, "******"},
	}
	for _, test := range tests {
		result := StringMask(test.params, test.star)
		assert.Equal(t, test.want, result)
	}

}

func TestWordSplit(t *testing.T) {
	words := "apple,bat,cat,goodbye,hello,yellow,why"

	tests := []struct {
		params [2]string
		want   string
	}{
		{[2]string{"hellocat", words}, "hello,cat"},
		{[2]string{"catbat", words}, "cat,bat"},
		{[2]string{"yellowapple", words}, "yellow,apple"},
		{[2]string{"", words}, "not possible"},
		{[2]string{"notcat", words}, "not possible"},
		{[2]string{"bootcamprocks!", words}, "not possible"},
	}
	for _, test := range tests {
		result := WordSplit(test.params)
		assert.Equal(t, test.want, result)
	}

}

func TestVariadicSet(t *testing.T) {
	tests := []struct {
		params []interface{}
		want   []interface{}
	}{
		{[]interface{}{4, 2, 5, 4, 2, 4}, []interface{}{4, 2, 5}},
		{[]interface{}{"bootcamp", "rocks!", "really", "rocks!"}, []interface{}{"bootcamp", "rocks!", "really"}},
		{[]interface{}{1, uint32(1), "first", 2, uint32(2), "second", 1, uint32(2), "first"}, []interface{}{1, uint32(1), "first", 2, uint32(2), "second"}},
	}

	for _, test := range tests {
		set := VariadicSet(test.params...)
		assert.Equal(t, test.want, set)
	}

}
