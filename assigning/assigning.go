package assigning // a new custom package assigning is created.
   import "fmt" 

func Add(){  // A new function Add is created and data will be given to it and execution would be done here.
	var x int // variable name,type is given but didn't declare the variable.
	var y int  // if we don't declare the variable it will take as "zero" for integer but it changes on datatype.
	x ,y = 54,63 // assigning value to the variables x and y.
	fmt.Println(x+y)//Value will be printed here by adding these two variables.
}

func Sub(){ // A new function Sub is created and data will be given to it and execution would be done here.
	var x int = 50  // variable name ,type are defined and declared the variable.
	var y int = 30
	fmt.Println(x-y)//value will be printed after subtraction by using format method.
}
func Mult(){
	var x = 20        //Only variable is declared but didn't assign the value .in this case variable take it as int as the decalred variable is int.
	 var y = 30 
	fmt.Println(x*y)
}

func Div(){ 
	x := 28  // shorthand declaration is used here ,no need to define type ,as it takes given number as int type.
	y := 56  
	fmt.Println(x/y)
}

