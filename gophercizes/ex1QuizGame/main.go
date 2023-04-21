package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)
var quizBucket = map[string]int{
	"Correct": 0,
	"Incorrect": 0,
	"Total": 0,
}

var c chan int

func main(){
	fmt.Println("You have 30 seconds to answer all questions")
  go timer()

  fileContent, err := os.Open("quiz.csv")
  if err != nil {
		fmt.Println("That didn't work!", err)
  }
  defer fileContent.Close()

	lines, err := csv.NewReader(fileContent).ReadAll()
	if err != nil {
		fmt.Println("Oops!", err)
	}
	quizBucket["Total"] = len(lines)
	for _, line := range lines {
			fmt.Println(line[0])
			var answer string
			fmt.Scanln(&answer)
		if answer == line[1] {
			fmt.Println("Yaaas!")
		  quizBucket["Correct"] = quizBucket["Correct"] + 1
		} else {
			fmt.Println("Sorry.")
			quizBucket["Incorrect"] = quizBucket["Incorrect"] + 1
		}
	}
  fmt.Printf("You got %d/%d correct, and %d/%d incorrect\n", quizBucket["Correct"], quizBucket["Total"], quizBucket["Incorrect"], quizBucket["Total"])
}

func timer(){
		
    time.Sleep(5 * time.Second)
		fmt.Println("You ran out of time!")
		os.Exit(0)
}