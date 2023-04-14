package main

import (
	"fmt"
	"net/http"
	"time"
)

func main () {
	 links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	 }

	 c := make(chan string) 

   for _, link := range	links {
		go checkLink(link, c) 
	 }

	//  for {
	// 	go checkLink(<-c, c)//<- GO knows that the channel returns a string, so it allows this first argument of a channel.
	//  } <= below is alternative syntax for the same thing here, for clarity

	for l := range c { //< this is saying wait for the channel to return some value, then run the body of hte for loop; we step through the for loop every time the channel emits some value. "Watch channel c for a value to come out, then assigne to 'l'"
	    go func(link string) {//this link string sets up the function with a type
				time.Sleep(time.Second * 5)//<= time.Second is an int64	
				checkLink(link, c)
			}(l) //<= this fires the function literal, so is where we pass the argument
			}
}

func checkLink(link string, c chan string ) {//"c is a channel that can only communicate strings"
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + "might be down!")
		c <- link
		return 
	}
	fmt.Println(link, "is up!")
	c <- link
}

