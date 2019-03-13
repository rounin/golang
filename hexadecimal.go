package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	go func() {
		for {
			fmt.Print(<-c)
		}
	}()
	for {
		select {
		case c <- "0":
		case c <- "1":
		case c <- "2":
		case c <- "3":
		case c <- "4":
		case c <- "5":
		case c <- "6":
		case c <- "7":
		case c <- "8":
		case c <- "9":
		case c <- "A":
		case c <- "B":
		case c <- "C":
		case c <- "D":
		case c <- "E":
		case c <- "F":
		}
	}
}
