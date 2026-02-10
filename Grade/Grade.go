package main

import "fmt"

func main() {
	var score float64
	fmt.Println("Please enter your score.(0-100): ")
	_, err := fmt.Scanln(&score)

	// Check if you,ve entered the numbers correcttly
	if err != nil {
		fmt.Println("lncorrect information: Please enter numbers only")
		return
	}

	// check if the score within the specified range.
	if score < 0 || score > 100 {
		fmt.Println("The score must be between 0 and 100.")
	} else if score >= 80 {
		fmt.Println(">grade is a A Excellent!")
	} else if score >= 70 {
		fmt.Println(">grade is a B Excellent!")
	} else if score >= 60 {
		fmt.Println(">grade is a C Excellent!")
	} else if score >= 50 {
		fmt.Println(">grade is a C Excellent!")
	} else if score >= 40 {
		fmt.Println(">grade is a D Excellent!")
	} else if score >= 30-20-10-5 {
		fmt.Println("Your grade re very poor and need to be improved!")
	}
}
