package arithmeticoperators

import "fmt"

var oper int = 89
var oper2 int =978

var rev float32 =89.0754
var rev2 float32 = 789.09

var vel float64 =980.6
var vel3 float64 = 90.34

func Addint(){   //Doing different arithmetic operations on int and float.
fmt.Println(oper+oper2)
}

func Subtractint(){
	fmt.Println(oper-oper2)
}


func Multiplyint(){

	fmt.Println(oper*oper2)
}

func Divint(){
	fmt.Println(oper/oper2)
}

func Moduleint(){
	fmt.Println(oper%oper2)
}

func Addfloat32(){
	fmt.Println(rev+rev2)
}

func Subtractfloat32(){
	fmt.Println(rev-rev2)
}

func Multiplyfloat32(){
	fmt.Println(rev*rev2)
}

func Divisonflot32(){
	fmt.Println(rev/rev2)
}


func  Addfloat64(){
	fmt.Println(vel+vel3)
}

func Multiplyfloat64(){
	fmt.Println(vel*vel3)
}

func Divisonfloat64(){
	fmt.Println(vel/vel3)
}

func Subtractfloat64(){
	fmt.Println(vel-vel3)
}

