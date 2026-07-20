package main

import "fmt"

func main() {
	// ==================
	// INTEGER TYPES
	// ==================

	var num1 int = 100     // Normal integer
	var num2 int8 = 127    // -128 to 127
	var num3 int16 = 32767 // -32768 to 32767
	var num4 int32 = 2147483647
	var num5 int64 = 9223372036854775807

	fmt.Println("num1:", num1)
	fmt.Println("num2:", num2)
	fmt.Println("num3:", num3)
	fmt.Println("num4:", num4)
	fmt.Println("num5:", num5)

	// ==================
	// FLOAT TYPES
	// ==================
	var f1 float32 = 3.14
	var f2 float64 = 3.141592653589793

	fmt.Println("\n=== FLOAT TYPES ===")
	fmt.Println("float32:", f1)
	fmt.Println("float64:", f2)

	// ==================
	// STRING TYPE
	// ==================
	var str1 string = "Hello Vikash"
	str2 := "Go is awesome!"

	fmt.Println("\n=== STRING TYPE ===")
	fmt.Println("str1:", str1)
	fmt.Println("str2:", str2)
	fmt.Println("Length:", len(str1))

	// ==================
	// BOOLEAN TYPE
	// ==================
	var b1 bool = true
	var b2 bool = false
	b3 := true

	fmt.Println("\n=== BOOLEAN TYPE ===")
	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)
	fmt.Println("b1 && b2:", b1 && b2) // AND
	fmt.Println("b1 || b2:", b1 || b2) // OR
	fmt.Println("!b1:", !b1)

	// ==================
	// BYTE AND RUNE
	// ==================
	var byteVal byte = 'A' // byte = uint8
	var runeVal rune = '😊' // rune = int32 (unicode)

	fmt.Println("\n=== BYTE AND RUNE ===")
	fmt.Println("byte:", byteVal)          // 65 (ASCII value)
	fmt.Printf("byte char: %c\n", byteVal) // A
	fmt.Println("rune:", runeVal)
	fmt.Printf("rune char: %c\n", runeVal) // // NOT
}
