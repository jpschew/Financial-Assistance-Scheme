package service

import (
	"FinancialAssistanceScheme/model"
	"errors"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Applicant = model.Applicant
type Household = model.HouseholdContent

func GetAllApplicants() ([]*Applicant, ServiceStatus) {

	a := &Applicant{}
	allApplicants, err := a.Get()
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, STATUS_DB_ERROR
		}
		return nil, STATUS_NO_APPLICANT_RECORD
	}

	return allApplicants, STATUS_OK
}

func CreateApplicant(firstName, lastName, nric string, employmentStatus, martialStatus, sex int, dob string, household []*Household) ServiceStatus {
	a := &Applicant{
		FirstName:        firstName,
		LastName:         lastName,
		NRIC:             nric,
		EmploymentStatus: employmentStatus,
		MartialStatus:    martialStatus,
		Sex:              sex,
		DateOfBirth:      dob,
		Household:        household,
	}

	if !validateNRIC(nric) {
		return STATUS_INVALID_NRIC
	}

	if !validateDOB(dob) {
		return STATUS_INVALID_DOB
	}

	for _, person := range a.Household {
		if !validateNRIC(person.NRIC) {
			return STATUS_INVALID_NRIC_HOUSEHOLD
		}

		if !validateDOB(person.DateOfBirth) {
			return STATUS_INVALID_DOB_HOUSEHOLD
		}

		if person.FirstName == "" || person.LastName == "" || person.DateOfBirth == "" || person.Relation == "" {
			return STATUS_PARAMS_EMPTY
		}
	}

	err := a.Create()
	if err != nil {
		return STATUS_DB_ERROR
	}

	return STATUS_OK
}

// validateNRIC validates a Singapore NRIC number
func validateNRIC(nric string) bool {
	if len(nric) != 9 {
		return false
	}

	// Extract prefix, number, and checksum character
	prefix := string(nric[0])
	number := nric[1:8]
	checksumChar := string(nric[8])

	// Validate prefix
	// S or T for singaporeans and PR
	// F or G for foreigners
	if !strings.Contains("STFG", prefix) {
		return false
	}

	// Convert prefix to its corresponding number
	prefixValue := map[string]int{
		"S": 0, "T": 4, "F": 0, "G": 4,
	}[prefix]

	// Weights for checksum calculation
	weights := []int{2, 7, 6, 5, 4, 3, 2}

	// Calculate weighted sum
	weightedSum := prefixValue
	for i, char := range number {
		digit := int(char - '0')
		weightedSum += digit * weights[i]
	}

	// Compute checksum
	checksum := weightedSum % 11
	checksumMap := "JZIHGFEDCBA" // starts with "S" or "T"
	if prefix == "F" || prefix == "G" {
		checksumMap = "XWUTRQPNMLK" // starts with "F" or "G"
	}
	calculatedChecksumChar := checksumMap[checksum]

	// Validate checksum character
	return checksumChar == string(calculatedChecksumChar)
}

const dobFormat = "2006-01-02" // Date format for YYYY-MM-DD

// validateDOB validates a date of birth in the format YYYY-MM-DD
func validateDOB(dob string) bool {
	// Parse the date using the defined format
	parsedDate, err := time.Parse(dobFormat, dob)
	if err != nil { // invalid date format, expected YYYY-MM-DD
		return false
	}

	// Ensure the parsed date is not in the future
	if parsedDate.After(time.Now()) { // date of birth cannot be in the future
		return false
	}

	return true
}
