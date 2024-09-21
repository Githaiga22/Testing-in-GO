// ----------------Table testing in GO-----------------------
package main

import "testing"

func TestAdd(t *testing.T) {
	//initialize a struct and its parameters as the testcases
	testcases := []struct{
		a, b, expected int
	} {
		{2, 2, 4},
		{0, 0, 0},
		{-2, -2, -4},
	}

	//range thru testcases to check each test case
	for _, tc := range testcases {
		//apply the add function parameters a and b and store them in result
		result := Add(tc.a, tc.b)
		//if result is not equal to expected return error
		if result != tc.expected {
			t.Errorf("Add(%d, %d) = %d; want %d;", tc.a, tc.b, result, tc.expected)
		}
	}
}