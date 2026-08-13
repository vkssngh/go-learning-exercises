package main

import "fmt"

// ============ PART 1 ============
func namaste() {
	fmt.Println("Namaste Go!")

}

// ============ PART 2 ============

func greet(name string) {
	fmt.Println("Hello", name)

}

// do parameters bhi ho sakte hain
func add(a int, b int) {
	result := a + b
	fmt.Println("Sum:", result)
}

// ============ PART 3 ============

// int matlab ye function ek number WAPAS dega
func multiply(a int, b int) int {
	return a * b
}

// string wapas dega
func getCity(code string) string {
	if code == "DEL" {
		return "Delhi"
	}
	return "Unknown City"
}

// ============ PART 4 ============

// Go mein ek se ZYADA value return ho sakti hai!
// Ye feature sirf Go mein hai - Java/Python mein nahi!

func divide(a int, b int) (int, string) {
	if b == 0 {
		return 0, "Error: Zero se divide nahi kar sakte!"
	}
	return a / b, "Success"
}

// Named return values - aur ek tarika
func rectangle(length int, width int) (area int, perimeter int) {
	area = length * width
	perimeter = 2 * (length + width)
	return // khud hi named values return karega
}

// ============ PART 5  ============
func getInfo() (string, int, string) {
	return "Vikash", 25, "Delhi"
}

// ============ PART 6 ============

// 1. Simple function
func showMenu() {
	fmt.Println("=== Calculator ===")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
}

// 2. Parameters wala
func addNumbers(a int, b int) int {
	return a + b
}

// 3. Multiple return
func calculate(a int, b int) (int, int, int) {
	sum := a + b
	diff := a - b
	product := a * b
	return sum, diff, product
}

// 4. Same type shorthand
func subtract(a, b int) int { // int ek baar likho
	return a - b
}

func main() {
	namaste() // Call function
	namaste() //
	namaste()
	greet("vikas")
	add(5, 10)

	result := multiply(4, 5)
	fmt.Println("4 x 5 =", result)

	city := getCity("INDIA")
	fmt.Println("City name is ", city)
	divResult, message := divide(10, 0)
	fmt.Println("10/0 answer", divResult, "Status", message)
	a, p := rectangle(5, 3)
	fmt.Println("Area:", a, "| Perimeter:", p)
	name, _, city := getInfo()
	fmt.Println("Name:", name, "| lives in  City:", city)

	showMenu()
	fmt.Println()

	fmt.Println("Addition:", addNumbers(10, 5))
	s, d, p := calculate(8, 3)
	fmt.Printf("Sum=%d, Diff=%d, Product=%d\n", s, d, p)

}
