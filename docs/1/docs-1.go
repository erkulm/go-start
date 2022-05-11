// Package dog offers functionalities that turns human years into dog years and vice versa
package dog

// ToHumanYears function takes an integer dogYears and returns the equivalent value in human years
func ToHumanYears(dogYears int) int {
	return dogYears * 7
}

// ToDogYears function takes an integer humanYears and returns the equivalent value in dog years
func ToDogYears(humanYears int) int {
	return humanYears / 7
}
