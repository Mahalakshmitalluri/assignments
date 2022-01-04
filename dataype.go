package main

import "fmt"



 func main(){
/* The value of variables are assigned with datatype,datatype int8 it can store 8bits of information,if that exceeds,it will show erroor and canno assign.
  Below examples are for different types of int8,int16,int32,int64 and the range of their storage.


*/
   var new int8 = 122 //Signed 8-bit integers (-128 to 127)
   var new1 int8 = 35
  // var c int8 =700 we cannot asign the value of 700 to c variable as int8 can only store -128 to 127.
   fmt.Println("The value of ",new-new1)
 

   var integer int16 = 3890//Signed 16-bit integers (-32768 to 32767)
   var integer1 int16 =7643
   // var g int16 = 50000 ,cannot assign this value ,as it also exceeds range limit of int16.
   fmt.Println("The value of ",integer-integer1)


   var cherry int32 = 678950//Signed 32-bit integers (-2147483648 to 2147483647)
   var veggies int32 = 908743
     fmt.Println("The value of ",cherry/veggies)
  
	 var int int64 =907432909875//Signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
     var int2 int64 =897654314
	 fmt.Println("The value of " ,int*int2)

     var may string = "macho" //string are used to store sequence of characters.
     var name string = "is something"
       fmt.Println(may+name)

	   var a bool= true // we only can assign either true or false ,if we are using boolean data type.
	   fmt.Println(a)
	   var b bool  = false
	   fmt.Println(b)
 }

 