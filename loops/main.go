package main

import "fmt"

// ============ PART 1 - Basic For Loop ============
func part1() {
	fmt.Println("===== PART 1 - Basic For Loop =====")

	// for loop ke 3 parts hain:
	// 1. i := 1      → shuruat kahan se
	// 2. i <= 5      → kab tak chalega
	// 3. i++         → har baar i mein 1 badhao

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}

// ============ PART 2 - For as While Loop ============
func part2() {
	fmt.Println("\n===== PART 2 - For as While Loop =====")

	// Go mein while nahi hota!
	// for ko while ki tarah use karte hain

	count := 1
	for count <= 5 {
		fmt.Println("Count:", count)
		count++ // count badhate raho
	}
}

// ============ PART 3 - Infinite Loop + Break ============
func part3() {
	fmt.Println("\n===== PART 3 - Infinite Loop + Break =====")

	// Infinite loop = hamesha chalta rahega
	// break se rokna padta hai!

	num := 1
	for {
		fmt.Println("Num:", num)
		num++
		if num > 5 {
			fmt.Println("Break! Ruk gaye")
			break // loop band karo
		}
	}
}

// ============ PART 4 - Continue ============
func part4() {
	fmt.Println("\n===== PART 4 - Continue =====")

	// continue = ye iteration skip karo
	// agle pe chale jao

	fmt.Println("Sirf Even Numbers:")
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue // odd hai toh skip karo
		}
		fmt.Println(i)
	}
}

// ============ PART 5 - Nested Loop ============
func part5() {
	fmt.Println("\n===== PART 5 - Nested Loop =====")

	// Loop ke andar Loop = Nested Loop

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
}

// ============ PART 6 - Range Loop ============
func part6() {
	fmt.Println("\n===== PART 6 - Range Loop =====")

	// Range = collection ke upar loop lagao
	// Bahut common hai Go mein!

	// String ke upar
	fruits := []string{"Apple", "Banana", "Mango", "Orange"}

	for index, fruit := range fruits {
		fmt.Println(index, "→", fruit)
	}

	fmt.Println()

	// Agar index nahi chahiye toh _ use karo
	for _, fruit := range fruits {
		fmt.Println("Fruit:", fruit)
	}
}

// ============ PART 7 - Real Life Example ============
func part7() {
	fmt.Println("\n===== PART 7 - Real Life Example =====")

	// Pahada (Table) banana
	num := 5
	fmt.Println(num, "Ka Pahada:")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", num, i, num*i)
	}

	fmt.Println()

	// 1 se 100 tak ka sum
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i // sum = sum + i
	}
	fmt.Println("1 se 100 tak ka Sum:", sum)
}

// ============ MAIN ============
func main() {
	part1()
	part2()
	part3()
	part4()
	part5()
	part6()
	part7()
}
