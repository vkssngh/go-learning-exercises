package main

import (
	"fmt"
)

// ============ PART 1 - Basic Array ============
func part1() {
	fmt.Println("===== PART 1 - Basic Array =====")

	// Array declare karne ka tarika:
	// var arrayNaam [size]dataType

	var marks [5]int // 5 size ka int array
	marks[0] = 85
	marks[1] = 90
	marks[2] = 78
	marks[3] = 92
	marks[4] = 88

	fmt.Println("Marks Array:", marks)
	fmt.Println("Marks[2]:", marks[2])         // 3rd element (index 2)
	fmt.Println("Lenght of array", len(marks)) // find the array length
}

// ============ PART 2 - Array Short Declaration ============
func part2() {
	fmt.Println("\n===== PART 2 - Array Short Declaration =====")

	// Seedha values deke banao
	fruits := [4]string{"Apple", "Banana", "Mango", "Orange"}
	fmt.Println("Fruits:", fruits)

	cities := [...]string{"Delhi", "Mumbai", "Bangalore"} // size auto calculate
	fmt.Println("City count:", len(cities))

	// loop on array
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])

	}
	// second way - range loop
	for i, fruit := range fruits {
		fmt.Println(i, "->", fruit)
	}
}

// ============ PART 3 - Array Limitation ============
func part3() {
	fmt.Println("\n===== PART 3 - Array Ki Problem =====")

	// Array ka size FIXED hai
	// Ek baar [3] diya toh sirf 3 hi rakhein ge

	nums := [3]int{10, 20, 30}
	fmt.Println("Array:", nums)
	fmt.Println("Size:", len(nums))

	// nums[3] = 40 → ERROR aayega! ❌
	// Isliye real code mein Slice use karte hain! ✅
	fmt.Println("Array ka size change nahi ho sakta!")
	fmt.Println("Isliye Slice use karte hain 👇")
}

// ============ PART 4 - Basic Slice ============
func part4() {
	fmt.Println("\n===== PART 4 - Basic Slice =====")

	// Slice mein [] mein size mat do!
	// Ye hi Array aur Slice ka fark hai

	var names []string // empty slice
	fmt.Println("Empty Slice:", names)
	fmt.Println("Length:", len(names))

	// Short way
	scores := []int{10, 20, 30, 40, 50}
	fmt.Println("Scores:", scores)
	fmt.Println("Length:", len(scores))

	// Element access karna
	fmt.Println("Pehla Score:", scores[0])
	fmt.Println("Aakhri Score:", scores[4])
}

// ============ PART 5 - Slice append ============
func part5() {
	fmt.Println("\n===== PART 5 - Slice mein add karna =====")

	// append se naye elements daalo!
	fruits := []string{"Apple", "Banana"}
	fmt.Println("Pehle:", fruits)

	// ek element add karo
	fruits = append(fruits, "Mango")
	fmt.Println("Mango add kiya:", fruits)

	// aur add karo
	fruits = append(fruits, "Orange", "Grapes")
	fmt.Println("Aur add kiya:", fruits)

	fmt.Println("Final Length:", len(fruits))
}

// ============ PART 6 - Slice Operations ============
func part6() {
	fmt.Println("\n===== PART 6 - Slice Operations =====")

	nums := []int{10, 20, 30, 40, 50}
	fmt.Println("Original:", nums)

	// Slicing - part nikalna
	// nums[start:end] → start se end-1 tak
	fmt.Println("nums[0:3] =", nums[0:3]) // 10,20,30
	fmt.Println("nums[2:5] =", nums[2:5]) // 30,40,50
	fmt.Println("nums[:3]  =", nums[:3])  // shuru se 3 tak
	fmt.Println("nums[2:]  =", nums[2:])  // 2 se aakhir tak
	fmt.Println("nums[:]   =", nums[:])   // sab kuch
}

// ============ PART 7 - make() function ============
func part7() {
	fmt.Println("\n===== PART 7 - make() se Slice =====")

	// make([]type, length, capacity)
	// jab pata ho kitna bada slice chahiye

	s := make([]int, 5) // 5 elements ka slice, sab 0
	fmt.Println("Make se:", s)
	fmt.Println("Length:", len(s))

	// Values assign karo
	s[0] = 100
	s[1] = 200
	s[2] = 300
	fmt.Println("Values daalne ke baad:", s)
}

// ============ PART 8 - Real Life Example ============
func part8() {
	fmt.Println("\n===== PART 8 - Real Life Example =====")

	// Student marks ka system
	studentMarks := []int{}

	// Marks add karo
	studentMarks = append(studentMarks, 85)
	studentMarks = append(studentMarks, 90)
	studentMarks = append(studentMarks, 78)
	studentMarks = append(studentMarks, 92)
	studentMarks = append(studentMarks, 88)

	fmt.Println("Saare Marks:", studentMarks)
	fmt.Println("Total Students:", len(studentMarks))

	// Total aur Average nikalo
	total := 0
	for _, mark := range studentMarks {
		total += mark
	}

	average := total / len(studentMarks)
	fmt.Println("Total Marks:", total)
	fmt.Println("Average Marks:", average)

	// Highest mark nikalo
	highest := studentMarks[0]
	for _, mark := range studentMarks {
		if mark > highest {
			highest = mark
		}
	}
	fmt.Println("Highest Mark:", highest)
}

func main() {

	part1()
	part2()
	part3()
	part4()
	part5()
	part6()
	part7()
	part8()
}
