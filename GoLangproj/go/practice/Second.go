package main

import "fmt"

//aprtof main package

func test() {
	fmt.Println("testing from Second")
}

func pyramid(){
	nr:=5
	rc:=1
	for i := nr; i >0; i-- {

		for j := 1; j <=i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <=rc; j++ {
			fmt.Print(rc," ")
		}
		fmt.Println()
		rc++
		
	}
	rc=nr-1
	for i := 1; i <=nr; i++ {
		for j := 0; j <=i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <=rc; j++ {
			fmt.Print(rc," ")
		}
		fmt.Println()
		rc--
	}
}