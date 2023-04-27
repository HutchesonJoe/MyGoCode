package main

import "fmt"

type Client struct {
	name string
	goal string
	weight int
}

var clients []*Client{}

func main() {
	var newName string
	fmt.Println("What is your client's name?")
	fmt.Scan(&newName)
	var newGoal string 
	fmt.Println("Goal?")
	fmt.Scan(&newGoal)
	var newWeight int 
	fmt.Println("Weight?")
	fmt.Scan(&newWeight)
  
	clients = clients.append("hello") 



}