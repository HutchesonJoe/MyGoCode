//
package main

import (
	"os"
	"fmt"
	"io/ioutil"
	
)

func main() {
	//file := os.Args[1]
	readFile, err := ioutil.ReadFile("./myfile.txt")
	if err != nil {
		fmt.Println("Oops, that didn't work:", err)
		os.Exit(1)
	}
  
	fmt.Printf("Contents: %s\n",readFile)

	//fmt.Println(io.Copy(os.Stdout, readFile))

}