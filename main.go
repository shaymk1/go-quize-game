package main

import (
	"fmt"
	
)



func main (){
	fmt.Println("Welcome to the game! ")

	var firstName string
	var age uint
	var match = "100%"

	fmt.Println("What is your first name?")
	fmt.Scan(&firstName)

	fmt.Printf("Hello, %v , are you ready to play today?\n" , firstName)

	fmt.Println("What is your age?")
    fmt.Scan(&age)

	if age >= 18 {
	fmt.Println("Yay! you are at the legal age to play")

}else {
	fmt.Println("You are too young to play")
	return
}

fmt.Printf(" Are you a morning person or a night person?    ")
var answer1, answer2 string
fmt.Scan(&answer1, &answer2)
fmt.Println(answer1, answer2)
if answer1 + " " + answer2 =="morning person" {

fmt.Println("Interesting person!")
}else{
	fmt.Printf("Congratulations! you are my %v match!Do you want to study together?" , match)
}





}