package main

import "fmt"

func part1() {
	fmt.Println("===== PART 1 - Basic Map =====")

	// Map banane ka tarika:
	// map[keyType]valueType

	// Way 1 - var se
	var person map[string]string
	fmt.Println("Empty Map:", person)

	// Way 2 - make se
	student := make(map[string]string)
	student["naam"] = "Vikash"
	student["city"] = "Delhi"
	student["course"] = "Go Language"

	fmt.Println("Student:", student)
	fmt.Println("Naam:", student["naam"])
	fmt.Println("City:", student["city"])
}

// ============ PART 2 - Short Declaration ============
func part2() {
	fmt.Println("\n===== PART 2 - Short Declaration =====")

	// Seedha values deke banao
	capitals := map[string]string{
		"India":   "New Delhi",
		"USA":     "Washington DC",
		"Japan":   "Tokyo",
		"England": "London",
	}

	fmt.Println("Capitals:", capitals)
	fmt.Println("India ki Capital:", capitals["India"])
	fmt.Println("Japan ki Capital:", capitals["Japan"])
}

// ============ PART 3 - Map Operations ============
func part3() {
	fmt.Println("\n===== PART 3 - Map Operations =====")

	scores := map[string]int{
		"Vikash": 90,
		"Ravi":   85,
		"Priya":  92,
	}

	fmt.Println("Pehle:", scores)

	// Add karo
	scores["Rahul"] = 88
	fmt.Println("Rahul add kiya:", scores)

	// Update karo
	scores["Ravi"] = 95
	fmt.Println("Ravi update kiya:", scores)

	// Delete karo
	delete(scores, "Vikash")
	fmt.Println("Vikash delete kiya:", scores)

	// Length
	fmt.Println("Total Students:", len(scores))
}

// ============ PART 4 - Key Check Karna ============
func part4() {
	fmt.Println("\n===== PART 4 - Key Exist Karta Hai? =====")

	// Bahut Important!
	// Agar key nahi hai toh Go error nahi deta
	// Default value deta hai - confusing ho sakta hai!

	student := map[string]int{
		"Maths":   90,
		"Science": 85,
		"English": 78,
	}

	// Wrong way - pata nahi chalega key hai ya nahi
	marks := student["Hindi"]          // key nahi hai!
	fmt.Println("Hindi Marks:", marks) // 0 aayega - confusing!

	// ✅ Sahi tarika - 2 values lo
	// value, ok := map[key]
	// ok = true  → key exist karti hai
	// ok = false → key nahi hai

	val, ok := student["Maths"]
	fmt.Println("\nMaths:", val, "| Exist?:", ok) // true

	val2, ok2 := student["Hindi"]
	fmt.Println("Hindi:", val2, "| Exist?:", ok2) // false

	// Practical use
	if marks, exists := student["Science"]; exists {
		fmt.Println("\nScience marks hain:", marks)
	} else {
		fmt.Println("\nScience marks nahi hain!")
	}
}

// ============ PART 5 - Map Loop ============
func part5() {
	fmt.Println("\n===== PART 5 - Map Loop =====")

	fruits := map[string]int{
		"Apple":  50,
		"Banana": 20,
		"Mango":  80,
		"Orange": 40,
	}

	// Key aur Value dono chahiye
	fmt.Println("Fruit Prices:")
	for fruit, price := range fruits {
		fmt.Printf("%-10s → ₹%d\n", fruit, price)
	}

	fmt.Println()

	// Sirf Keys chahiye
	fmt.Println("Sirf Fruits:")
	for fruit := range fruits {
		fmt.Println("→", fruit)
	}
}

// ============ PART 6 - Nested Map ============
func part6() {
	fmt.Println("\n===== PART 6 - Nested Map =====")

	// Map ke andar Map!
	// Real life - Student ka poora record

	students := map[string]map[string]int{
		"Vikash": {
			"Maths":   90,
			"Science": 85,
			"English": 78,
		},
		"Ravi": {
			"Maths":   75,
			"Science": 80,
			"English": 90,
		},
	}

	fmt.Println("Vikash ke Marks:")
	for subject, marks := range students["Vikash"] {
		fmt.Printf("  %-10s: %d\n", subject, marks)
	}

	fmt.Println("\nRavi ke Maths Marks:", students["Ravi"]["Maths"])
}

// ============ PART 7 - Real Life Example ============
func part7() {
	fmt.Println("\n===== PART 7 - Real Life Example =====")

	// Word Counter - Sentence mein words count karo
	sentence := []string{"go", "is", "fast", "go", "is", "easy", "go"}

	wordCount := make(map[string]int)

	for _, word := range sentence {
		wordCount[word]++ // har baar word aaye - count badhao
	}

	fmt.Println("Word Count:", wordCount)
	fmt.Println("Words and their counts individually:")
	for word, count := range wordCount {
		fmt.Printf("  %-5s: %d\n", word, count)
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
