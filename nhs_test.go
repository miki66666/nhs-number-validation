package nhs

import (
	"testing"
)

func TestCheckNHS(t *testing.T) {
	cases := []struct {
		nhsNumber string
		expected  bool
	}{
		{"5990128088", true},
		{"1275988113", true},
		{"4536026665", true},
		{"5990128087", false},
		{"4536016660", false},
	}

	for _, tc := range cases {
		t.Run(tc.nhsNumber, func(t *testing.T) {
			result := CheckNHSNumber(tc.nhsNumber)
			if result != tc.expected {
				t.Errorf("The NHS Number %s, was expected %v, but instead got %v", tc.nhsNumber, tc.expected, result)
			}
		})
	}
}

func TestGenerateValidhsNHSNumber(t *testing.T) {
	nhsNumber := GenerateValidNHSNumber()

	if !CheckNHSNumber(nhsNumber) {
		t.Errorf("Generated NHS Number %s is not valid", nhsNumber)
	}
}
