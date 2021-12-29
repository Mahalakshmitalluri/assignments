package variable // creating a custom package 

import "fmt" 

func Add(){ // creating a function Add to add two integer values.function name couldn't effect the assigning 
	var x int = 4 //Declaring values to x and y variables with type name 
	var y int = 20
	fmt.Println(x+y)//used the impport format to print the variables by adding
}

func Subtract(){// creating a function sub to sub two integer values.
	 var x = 25     //assigning  values to x and y variables , the type will be recognised as int as its given integer value.
	 var y = 30
	 fmt.Println(x-y)
}

func Multiple(){  //creating a function multiple to multiply two integer values.
	x := 560       //Declaring values to x and y variables with short hand declaration which takes it as variable and type it is with data declared which is integer here.
	y := 52
	fmt.Println(x*y)
}
func Div(){  //creating a function div to divide two integer values.
	var x int// variable name ,type will be given in this syntax
	var y int 
	x = 2  // value to variable is aassigned here , if we don't assign any values to variable it will take as zero for integer it defers to different types like string cd
    y =10
	fmt.Println(x/y)
}