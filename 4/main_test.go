package main

import "testing"

func TestPassport_ValidateByr(t *testing.T) {
	type TestCase struct {
		Input Passport
		Expected bool
	}

	testCases := []TestCase{
		{
			Input: Passport{
				Byr: "2002",
			},
			Expected: true,
		},
		{
			Input: Passport{
				Byr: "2003",
			},
			Expected: false,
		},
		{
			Input: Passport{
				Byr: "200",
			},
			Expected: false,
		},
		{
			Input: Passport{
				Byr: "20l0",
			},
			Expected: false,
		},
	}

	for _, testCase := range testCases {
		if testCase.Input.ValidateByr() != testCase.Expected {
			t.Errorf("Testing value %s, expetcted %t, got %t", testCase.Input.Byr, testCase.Expected, testCase.Input.ValidateByr())
		}
	}
}

func TestPassport_ValidateHgt(t *testing.T) {
	type TestCase struct {
		Input Passport
		Expected bool
	}

	testCases := []TestCase{
		{
			Input: Passport{
				Hgt: "60in",
			},
			Expected: true,
		},
		{
			Input: Passport{
				Hgt: "190cm",
			},
			Expected: true,
		},
		{
			Input: Passport{
				Hgt: "190in",
			},
			Expected: false,
		},
		{
			Input: Passport{
				Hgt: "190",
			},
			Expected: false,
		},
	}

	for _, testCase := range testCases {
		if testCase.Input.ValidateHgt() != testCase.Expected {
			t.Errorf("Testing value %s, expetcted %t, got %t", testCase.Input.Hgt, testCase.Expected, testCase.Input.ValidateHgt())
		}
	}
}