package main

import "fmt"

func main() {
    var age int
    fmt.Print("Enter your age: ")
    fmt.Scan(&age)
    fmt.Printf("Hello, World! I am %d years old.\n", age)
}