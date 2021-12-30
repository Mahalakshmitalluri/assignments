package assigning // a new custom package assigning is created.
   import "fmt" 

func Add(){  // A new function Add is created and data will be given to it and execution would be done here.
	var x int // variable name,type is given but didn't declare the variable.
	var y int  // if we don't declare the variable it will take as "zero" for integer but it changes on datatype.
	x ,y = 54,63 // assigning value to the variables x and y.
	fmt.Println("The value after adding  is ",x+y)//Value will be printed here by adding these two variables.
}

func Sub(){ // A new function Sub is created and data will be given to it and execution would be done here.
	var x int = 50  // variable name ,type are defined and declared the variable.
	var y int = 30
	fmt.Println("The value after subtracting  is ",x-y)//value will be printed after subtraction by using format method.
}
func Mult(){
	var x = 20        //Only variable is declared but didn't assign the value .in this case variable take it as int as the decalred variable is int.
	 var y = 30 
	fmt.Println("The value after mulptilying is ",x*y)
	
}

func Div(){ 
	x := 28  // shorthand declaration is used here ,no need to define type ,as it takes given number as int type.
	y := 25  
	fmt.Println("The value after dividing  is ",x/y)
}

func Person1(){
	var name string ="sasi"//variables are declared with a string name and integer value with defining the type.
	var age int     = 21
	fmt.Println("Name and age of person1 is ",name,"-",age)
}

func Person2(){
	var name = "manju"//variables are declared with a string name and integer value but type is not defined but it  takes name and value and it recognises the type.
	var age = 23
	fmt.Println("Name and age of person2 is ",name,"-",age)
}

func Person3(){
	name := "nandy" //  short hand declaration is ":=" is used to declare name and age and also takes the assigned value as integer or string.
	age := 25   
	fmt.Println("Name and age of person3 is ",name,"-",age)
}

func Person4(){  //function Person are created in here.
	var name string //name and string are taken without declaring , string executes a blank value and int takes a zero value.
	var age int 
	name = "manu"
	age = 29
 fmt.Println("Name and age of person4 is ",name,"-",age)
}