package main

import "fmt"

func main() {
	// ==================
	// WAY 1: var keyword
	// ==================
	var name string = "vikash"
	var age int = 35
	var isMarried bool = true
	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("isMarried:", isMarried)

	// WAY 2: Short Declaration
	city := "Mumbai"
	pincode := 400001
	temperature := 36.5
	fmt.Println("city:", city, "pincode:", pincode, "temperature:", temperature)
	// =======================
	// WAY 3: Multiple variables
	// =======================
	var (
		country  string = "India"
		language string = "Go"
		year     int    = 2024
	)
	fmt.Println("country:", country, "language:", language, "year:", year)
	var (
		firstName = "Vikash"
		lastName  = "Singh"
		countryy  = "India"
	)

	fmt.Println("\n=== WAY 3: Multiple var ===")
	fmt.Println("Full Name:", firstName, lastName)
	fmt.Println("Country:", countryy)

	// =======================
	// WAY 4: Zero Values
	// Default values in Go
	// =======================
	var defaultInt int
	var defaultString string
	var defaultBool bool
	var defaultFloat float64
	fmt.Println("\n=== WAY 4: Zero/Default Values ===")
	fmt.Println("Default int:", defaultInt)       // 0
	fmt.Println("Default string:", defaultString) // ""
	fmt.Println("Default bool:", defaultBool)     // false
	fmt.Println("Default float:", defaultFloat)   // 0

	// =======================
	// WAY 5: Type Conversion
	// =======================
	var intNum int = 42
	var floatNum float64 = float64(intNum) // int to float
	var backToInt int = int(floatNum)      // float to int

	fmt.Println("\n=== WAY 5: Type Conversion ===")
	fmt.Println("Int:", intNum)
	fmt.Println("Float:", floatNum)
	fmt.Println("Back to Int:", backToInt)
}
