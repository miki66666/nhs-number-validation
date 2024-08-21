package nhs

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// GenerateValidNHSNumber generates a valid NHS number
func GenerateValidNHSNumber() string {
	rand.Seed(time.Now().UnixNano())

	for {
		// Generate a random nine digit number including zeros
		nineDigits := fmt.Sprintf("%09d", rand.Intn(1000000000))

		// Get the NHS number and check if it is valid
		digit, err := getLastExpectedDigit(nineDigits)
		if err == nil && digit != 10 {
			return nineDigits + strconv.Itoa(digit)
		}
	}
}

// getLastExpectedDigit calculates the last expected digit for the NHS number
func getLastExpectedDigit(nhsNumber string) (int, error) {
	// Define the mul from 10 and decrementing
	mul := 10
	var sum int

	for i := 0; i < 9; i++ {
		digit, err := strconv.Atoi(string(nhsNumber[i]))
		if err != nil {
			return 0, err
		}
		sum += mul * digit
		mul--
	}

	remainder := sum % 11

	// Edge case of remainder 0
	if remainder == 0 {
		return 0, nil
	}
	// Edge case of remainder 1
	if remainder == 1 {
		return 0, errors.New("invalid NHS number with remainder 1")
	}
	checkDigit := 11 - remainder
	return checkDigit, nil
}

// CheckNHSNumber validates an NHS number
func CheckNHSNumber(nhsNumber string) bool {
	if len(nhsNumber) != 10 {
		return false
	}

	expectedDigit, err := getLastExpectedDigit(nhsNumber[:9])
	if err != nil {
		return false
	}

	lastDigit, err := strconv.Atoi(string(nhsNumber[9]))
	if err != nil {
		return false
	}

	return expectedDigit == lastDigit
}
