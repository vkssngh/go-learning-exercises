package main

import "fmt"

const Appname = "My Go App"
const Version = "1.0.0"
const MaxUsers = 1000

const (
	StatusPending  = iota // 0
	StatusActive          // 1
	StatusInactive        // 2
	StatusDeleted         // 3
)

// Days of week
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {

	fmt.Println("=== BASIC CONSTANTS ===")
	fmt.Println("App:", Appname)
	fmt.Println("Version:", Version)
	fmt.Println("Max Users:", MaxUsers)

	// Iota - Status
	fmt.Println("\n=== STATUS (iota) ===")
	fmt.Println("Pending:", StatusPending)   // 0
	fmt.Println("Active:", StatusActive)     // 1
	fmt.Println("Inactive:", StatusInactive) // 2
	fmt.Println("Deleted:", StatusDeleted)   // 3

	// Iota - Days
	fmt.Println("\n=== DAYS (iota) ===")
	fmt.Println("Sunday:", Sunday)       // 0
	fmt.Println("Monday:", Monday)       // 1
	fmt.Println("Wednesday:", Wednesday) // 3
	fmt.Println("Saturday:", Saturday)   // 6

	// Constants ka type
	const Pi = 3.14159
	const Greeting = "Hello"
	const IsReady = true

	fmt.Println("\n=== CONSTANTS IN FUNCTION ===")
	fmt.Println("Pi:", Pi)
	fmt.Println("Greeting:", Greeting)
	fmt.Println("IsReady:", IsReady)
	// Constant use case - real example
	const MaxLoginAttempts = 3
	attempts := 0

	fmt.Println("\n=== REAL USE CASE ===")
	fmt.Println("Max attempts allowed:", MaxLoginAttempts)
	fmt.Println("Current attempts:", attempts)
	fmt.Println("Can try again?", attempts < MaxLoginAttempts)

}
