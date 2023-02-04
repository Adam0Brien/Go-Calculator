package main

import (
	"fmt"
	"math"
	"bufio"
	"os"
	"strconv"
	"strings"
)


func main(){
	fmt.Println("Calculator App")
	menu()
}

func menu(){
	
	fmt.Print("\n 1): Add\n 2): Subtract\n 3): Multiply\n 4): Divide\n 5): Exit\n")
	var i int
	fmt.Scan(&i)
    //fmt.Print("Write ", i, " as ")
	fmt.Println(i)

    switch i {
		case 1:
			add()
		case 2:
			subtract()
		case 3:
			multiply()
		case 4:
			divide()
		case 5:
			fmt.Println("Goodbye!")
		default:
			fmt.Print("Choose an Int between 1 and 4\n")
			menu()
	}
}

func add(){

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Add Two Integers Together. ")
	fmt.Print("Enter first integer: ")
	input1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value must be a number")
	}
	
	

	fmt.Print("Enter second integer: ")
	input2, _ := reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value must be a number")
	}
	 

	
	sum := float1 + float2


	fmt.Print(float1 ,"+", float2,"=",math.Round(sum*100) / 100)
	


	menu()
}

func subtract(){

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Subtract one number from another. ")
	fmt.Print("Enter first integer: ")
	input1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value must be a number")
	}
	
	

	fmt.Print("Enter second integer: ")
	input2, _ := reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value must be a number")
	}
	 

	
	sum := float1 - float2

	fmt.Print(float1 ,"-", float2,"=",math.Round(sum*100) / 100)

	menu()
}


func multiply(){

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Multiply one number to another. ")
	fmt.Print("Enter first integer: ")
	input1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value must be a number")
	}
	
	

	fmt.Print("Enter second integer: ")
	input2, _ := reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value must be a number")
	}
	 

	
	sum := float1 * float2

	fmt.Print(float1 ,"x", float2,"=",math.Round(sum*100) / 100)

	menu()
}


func divide(){

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("divide one number from another. ")
	fmt.Print("Enter first integer: ")
	input1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value must be a number")
	}
	
	

	fmt.Print("Enter second integer: ")
	input2, _ := reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value must be a number")
	}
	 

	
	sum := float1 / float2

	fmt.Print(float1 ,"/", float2,"=",math.Round(sum*100) / 100)

	menu()
}
