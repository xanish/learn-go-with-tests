package _5_intro_to_property_based_tests

import "strings"

// RomanNumeral represents a mapping between an integer value and its Roman numeral symbol.
type RomanNumeral struct {
	Value  uint16
	Symbol string
}

// RomanNumerals is a predefined slice of RomanNumeral structs that maps integers to Roman numeral symbols.
var RomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ConvertToRoman converts an integer (arabic) into its Roman numeral representation.
func ConvertToRoman(arabic uint16) string {
	// Initialize a strings.Builder to efficiently construct the result.
	var result strings.Builder

	for _, numeral := range RomanNumerals {
		// Append the Roman numeral symbol to the result repeatedly till the arabic number is
		// greater than or equal to the current numeral's value.
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) uint16 {
	var arabic uint16 = 0

	for _, numeral := range RomanNumerals {
		// Continue adding the numeral's value to arabic while the roman string
		// starts with the current numeral's symbol.
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}

	return arabic
}
