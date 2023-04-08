package main

import (
	"fmt"
)

type clientStats struct {
	weight int
	goal string
}

type client struct {
	firstName string
	lastName 	string
	email 		string
	stats 		clientStats
}
type yogaClient client

func main() {
	emily := client{
		firstName: "Emily",
		lastName: "Madison",
		email: "emad@gmail.com",
		stats: clientStats {86, "cardio"},
	}
	aparnaa := yogaClient{
		firstName: "Aparnaa",
	}
	fmt.Println(emily)
	fmt.Println(aparnaa)
}