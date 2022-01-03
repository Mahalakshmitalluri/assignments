package scopes // created anotherfile under scopes package with packagescope name.

import"fmt"                     // variable created in one package canot accessible in other package.
       
//packagescope 
func Example3(){      // created a function here and executing it.
	vscode := "string"
	fmt.Println(vscode)
	fmt.Println(Terrace)// variable "terrace" is taken from another  file  filescopes.go and can exeute here too, as they are under samepackage called "scopes".
}


/* packagelevel scope can take variables inside it ,not from inside the functions 
,but declared outside the function under the same package level
*/
