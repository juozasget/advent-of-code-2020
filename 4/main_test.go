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

func TestPassport_ValidateHcl(t *testing.T) {
	type TestCase struct {
		Input Passport
		Expected bool
	}

	testCases := []TestCase{
		{
			Input: Passport{
				Hcl: "#888785",
			},
			Expected: true,
		},
		{
			Input: Passport{
				Hcl: "123abc",
			},
			Expected: false,
		},
		{
			Input: Passport{
				Hcl: "#123abz",
			},
			Expected: false,
		},
	}

	for _, testCase := range testCases {
		if testCase.Input.ValidateHcl() != testCase.Expected {
			t.Errorf("Testing value %s, expetcted %t, got %t", testCase.Input.Hcl, testCase.Expected, testCase.Input.ValidateHcl())
		}
	}
}

func TestPassport_ValidateAll(t *testing.T) {
	type TestCase struct {
		Input Passport
		Expected bool
	}

	testCases := []TestCase{
		{
			Input: Passport{"1989", "2013", "2022", "155cm","#733820", "grn", "728471979", "someId", true},
			Expected: true,
		},
	}

	for _, testCase := range testCases {
		if !testCase.Input.ValidateByr() {
			t.Errorf("Testing value Byr: %s, expetcted %t, got %t", testCase.Input.Byr, testCase.Expected, testCase.Input.ValidateByr())
		}
		if !testCase.Input.ValidateHgt() {
			t.Errorf("Testing value Hgt: %s, expetcted %t, got %t", testCase.Input.Hgt, testCase.Expected, testCase.Input.ValidateHgt())
		}
		if !testCase.Input.ValidatePid() {
			t.Errorf("Testing value Pid: %s, expetcted %t, got %t", testCase.Input.Pid, testCase.Expected, testCase.Input.ValidatePid())
		}
		if !testCase.Input.ValidateIyr() {
			t.Errorf("Testing value Iyr: %s, expetcted %t, got %t", testCase.Input.Iyr, testCase.Expected, testCase.Input.ValidateIyr())
		}
		if !testCase.Input.ValidateEcl() {
			t.Errorf("Testing value Ecl: %s, expetcted %t, got %t", testCase.Input.Ecl, testCase.Expected, testCase.Input.ValidateEcl())
		}
		if !testCase.Input.ValidateHcl() {
			t.Errorf("Testing value Hcl: %s, expetcted %t, got %t", testCase.Input.Hcl, testCase.Expected, testCase.Input.ValidateHcl())
		}
		if !testCase.Input.ValidateEyr() {
			t.Errorf("Testing value Eyr: %s, expetcted %t, got %t", testCase.Input.Eyr, testCase.Expected, testCase.Input.ValidateEyr())
		}
	}
}


