package main

import "fmt"

// ============ PART 1 - Basic if/else ============

// ============ PART 1 - Basic if/else ============
func part1() {
	fmt.Println("===== PART 1 - Basic if/else =====")

	age := 18

	if age >= 18 {
		fmt.Println("Tum adult ho ✅")
	} else {
		fmt.Println("Tum minor ho ❌")
	}

	// Sirf if - koi else nahi
	marks := 85
	if marks > 80 {
		fmt.Println("Distinction milega! 🎉")
	}
}

// ============ PART 2 - if / else if / else ============
func part2() {
	fmt.Println("\n===== PART 2 - if/else if/else =====")

	marks := 72

	if marks >= 90 {
		fmt.Println("Grade: A+")
	} else if marks >= 80 {
		fmt.Println("Grade: A")
	} else if marks >= 70 {
		fmt.Println("Grade: B")
	} else if marks >= 60 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: Fail ❌")
	}
}

// ============ PART 3 - if with short statement ============
func part3() {
	fmt.Println("\n===== PART 3 - if Short Statement =====")

	// Go ki specialty!
	// Variable declare karo AUR check karo - ek line mein!
	if num := 42; num%2 == 0 {
		fmt.Println(num, "Even number hai ✅")
	} else {
		fmt.Println(num, "Odd number hai")
	}
	// Note: num sirf if/else ke andar kaam karega bahar nahi!
}

// ============ PART 4 - Basic Switch ============
func part4() {
	fmt.Println("\n===== PART 4 - Basic Switch =====")

	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Aaj Monday hai - Office jaao! 😫")
	case "Saturday":
		fmt.Println("Aaj Saturday hai - Aaram karo! 😎")
	case "Sunday":
		fmt.Println("Aaj Sunday hai - Movie dekho! 🎬")
	default:
		fmt.Println("Normal weekday hai")
	}
}

// ============ PART 5 - Switch Multiple Cases ============
func part5() {
	fmt.Println("\n===== PART 5 - Switch Multiple Cases =====")

	day := "Saturday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday hai - Kaam karo! 💼")
	case "Saturday", "Sunday":
		fmt.Println("Weekend hai - Enjoy karo! 🎉")
	default:
		fmt.Println("Invalid day!")
	}
}

// ============ PART 6 - Switch without condition ============
func part6() {
	fmt.Println("\n===== PART 6 - Switch without Condition =====")

	// if/else ki tarah use kar sakte ho switch ko!
	marks := 85

	switch {
	case marks >= 90:
		fmt.Println("Grade: A+")
	case marks >= 80:
		fmt.Println("Grade: A")
	case marks >= 70:
		fmt.Println("Grade: B")
	default:
		fmt.Println("Grade: C ya Fail")
	}
}

// ============ PART 7 - Real Life Example ============
func part7() {
	fmt.Println("\n===== PART 7 - Real Life Example =====")

	salary := 50000
	experience := 3

	// Multiple conditions - && matlab AND, || matlab OR
	if salary > 40000 && experience > 2 {
		fmt.Println("Loan approved! ✅")
	} else if salary > 40000 || experience > 5 {
		fmt.Println("Loan review mein hai 🔄")
	} else {
		fmt.Println("Loan rejected ❌")
	}
}
func main() {
	part1()
	part2()
	part3()
	part4()
	part5()
	part6()
	part7()

}
