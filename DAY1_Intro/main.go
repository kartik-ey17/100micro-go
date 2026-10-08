package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Println("\n+=+=+= Go greeter =+=+=+")
	fmt.Print("What's your name? : ")
	fmt.Scan(&name)
	fmt.Print("What's your age? : ")
	fmt.Scan(&age)
	diff := age * -1
	fmt.Printf("\nHellow %s!\nSo your %d years old.", name, age)
	if age > 60 {
		fmt.Printf("\nYou are a senior citizen!")
	} else if age > 18 {
		fmt.Printf("\nYou are an adult.")
	} else if age > 0 {
		fmt.Print("Your a kid!")
	} else {
		fmt.Printf("\n%d years for you to come to this planet son", diff)
	}
	fmt.Println("\nWelcome to GoLang.")
}
