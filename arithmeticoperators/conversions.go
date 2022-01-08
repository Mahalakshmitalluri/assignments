package arithmeticoperators
import "fmt"

var a int8 =98
var b int16 =876
var c int32 =78
var d int64 = 789
var e float32 =87.76
var f float64 =908.6

func Mix(){

	t :=  int16(a)+(b)
	fmt.Println("the result is ",t)
}

func New(){
	v := int32(b)+c//int16 is converted into int32.
	fmt.Println("the result is ",v)
}
 
func Elf(){
	r := int64(c)+d
	fmt.Println("the result is ",r)
}

func Elsa(){
	h := int64(c)+int64(e)//float32 value of e is converted into integer value
	fmt.Println("the result is ",h)
}

func Snow(){
	j := float64(e)+f//float32 value of e is converted to float64,but we can't convert float64 into float32. As , float32 storage is less from float64.
	fmt.Println(j)//basically we can't go from higher storage to low storage convertions.
}