package scopes
import "fmt"

  var Terrace string = "mall"
func Example(){
	var x int = 6
	numberone := 34
	numbertwo := 89
fmt.Println(numberone-numbertwo,x,Terrace)
}
/* we cannot access var x int from examplefunction into example2function
it is the functional scope which is locally created variables cannot acces from another function

*/
func Example2(){
	var z int =87
	fmt.Println(z,Terrace)
	//fmt.Println(x)we cannot do this approach as varable x int is locally created in examplefunction.
}