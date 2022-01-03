package main

import (
	"fmt"

	"github.com/Mahalakshmitalluri/assignments/globalscope"
	"github.com/Mahalakshmitalluri/assignments/scopes"
)
func main(){
	scopes.Example()
	scopes.Example3()
	fmt.Println(scopes.Terrace)// we can access the varaible declared inside the package/functions to outside of package can be accessible only when we use uppercase letter.
	globalscope.New()
	//In scopes package two functions are created but we don't need to take both of the functions to execution if not required.
	//we also can't print the var x value in examplefunction sepertely here. 
//fmt.Println(scopes.example(x))
}

