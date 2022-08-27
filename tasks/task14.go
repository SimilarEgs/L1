package main

import "fmt"

//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

func main() {

	varInt := 5
	varString := "Hello World"
	VarBool := true
	varChanInt := make(chan int, 0)

	typeChecker(varInt)
	typeChecker(varString)
	typeChecker(VarBool)
	typeChecker(varChanInt)
}

func typeChecker(inputVar interface{}) {

	switch varType := inputVar.(type) {
	case int:
		fmt.Printf("[Info] type of the variable - %T\n", varType)
	case string:
		fmt.Printf("[Info] type of the variable - %T\n", varType)
	case bool:
		fmt.Printf("[Info] type of the variable - %T\n", varType)
	case chan int:
		fmt.Printf("[Info] type of the variable - %T\n", varType)
	}

}
