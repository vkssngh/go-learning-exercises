package main

import "fmt"

// ============ STRUCT DEFINE KARO ============
// Ye main() ke bahar likhte hain!

// Basic Struct
type Person struct {
	naam string
	age  int
	city string
}

// Student Struct
type Student struct {
	naam    string
	age     int
	marks   int
	passing bool
}

// Address Struct
type Address struct {
	street  string
	city    string
	state   string
	pincode int
}

// Nested Struct - Struct ke andar Struct!
type Employee struct {
	naam    string
	age     int
	salary  float64
	address Address // Address struct andar!
}

// ============ PART 1 - Basic Struct ============
func part1() {
	fmt.Println("===== PART 1 - Basic Struct =====")

	// Way 1 - var se banao
	var p1 Person
	p1.naam = "Vikash"
	p1.age = 25
	p1.city = "Delhi"

	fmt.Println("Person:", p1)
	fmt.Println("Naam:", p1.naam)
	fmt.Println("Age:", p1.age)
	fmt.Println("City:", p1.city)
}

// ============ PART 2 - Short Declaration ============
func part2() {
	fmt.Println("\n===== PART 2 - Short Declaration =====")

	// Way 2 - seedha values deke banao
	p2 := Person{
		naam: "Ravi",
		age:  30,
		city: "Mumbai",
	}

	fmt.Println("Person:", p2)

	// Way 3 - order mein values deke banao
	// (field names nahi likhte - order same rakhna hoga!)
	p3 := Person{"Priya", 28, "Bangalore"}
	fmt.Println("Person:", p3)
}

// ============ PART 3 - Struct Update Karna ============
func part3() {
	fmt.Println("\n===== PART 3 - Struct Update =====")

	p := Person{
		naam: "Vikash",
		age:  25,
		city: "Delhi",
	}

	fmt.Println("Pehle:", p)

	// Update karo
	p.age = 26
	p.city = "Mumbai"

	fmt.Println("Baad mein:", p)
}

// ============ PART 4 - Struct with Function ============
func part4() {
	fmt.Println("\n===== PART 4 - Struct with Function =====")

	// Function mein struct pass karo
	s1 := Student{
		naam:    "Vikash",
		age:     20,
		marks:   85,
		passing: true,
	}

	s2 := Student{
		naam:    "Ravi",
		age:     21,
		marks:   35,
		passing: false,
	}

	printStudent(s1)
	printStudent(s2)
}

// Student print karne ka function
func printStudent(s Student) {
	fmt.Printf("\nStudent: %s\n", s.naam)
	fmt.Printf("Age: %d\n", s.age)
	fmt.Printf("Marks: %d\n", s.marks)
	if s.passing {
		fmt.Println("Result: PASS ✅")
	} else {
		fmt.Println("Result: FAIL ❌")
	}
}

// ============ PART 5 - Nested Struct ============
func part5() {
	fmt.Println("\n===== PART 5 - Nested Struct =====")

	emp := Employee{
		naam:   "Vikash",
		age:    25,
		salary: 50000.00,
		address: Address{
			street:  "MG Road",
			city:    "Delhi",
			state:   "Delhi",
			pincode: 110001,
		},
	}

	fmt.Println("Employee:", emp.naam)
	fmt.Println("Salary:", emp.salary)
	fmt.Println("City:", emp.address.city)       // nested access
	fmt.Println("Pincode:", emp.address.pincode) // nested access
}

func part6() {
	fmt.Println("\n===== PART 6 - Slice of Structs =====")

	// Real life mein kai log hote hain!
	// Slice mein structs rakho

	persons := []Person{
		{naam: "Vikash", age: 25, city: "Delhi"},
		{naam: "Ravi", age: 30, city: "Mumbai"},
		{naam: "Priya", age: 28, city: "Bangalore"},
		{naam: "Rahul", age: 22, city: "Chennai"},
	}

	fmt.Println("Saare Persons:")
	for i, p := range persons {
		fmt.Printf("%d. %-10s | Age: %d | City: %s\n",
			i+1, p.naam, p.age, p.city)
	}

	// Sabse bada age dhundho
	oldest := persons[0]
	for _, p := range persons {
		if p.age > oldest.age {
			oldest = p
		}
	}
	fmt.Println("\nSabse Bade:", oldest.naam, "| Age:", oldest.age)
}

// ============ PART 7 - Real Life Example ============
func part7() {
	fmt.Println("\n===== PART 7 - Real Life Example =====")

	// Product management system
	type Product struct {
		id    int
		naam  string
		price float64
		stock int
	}

	products := []Product{
		{id: 1, naam: "Laptop", price: 50000, stock: 10},
		{id: 2, naam: "Phone", price: 15000, stock: 25},
		{id: 3, naam: "Tablet", price: 20000, stock: 5},
	}

	fmt.Println("=== Product List ===")
	for _, p := range products {
		fmt.Printf("ID: %d | %-10s | ₹%.2f | Stock: %d\n",
			p.id, p.naam, p.price, p.stock)
	}

	// Total inventory value
	totalValue := 0.0
	for _, p := range products {
		totalValue += p.price * float64(p.stock)
	}
	fmt.Printf("\nTotal Inventory Value: ₹%.2f\n", totalValue)
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
