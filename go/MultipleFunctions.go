package main
import "fmt"

import "calculation"

func mynewfunc(a int,b int)(int){
	c:=a+b
	return c
}

func main(){
	
x := 10
	fmt.Println("this the value of x ",x)
	y:=x+x
	a:="Stonks"
	const aa = "Vegtals"
	fmt.Println("this the value of y, a and aa",y,a,aa)
	var i int
	for i=0; i<5; i++ {
		for j:=0;j<(i+1);j++{
			fmt.Print("Stonks ")
		}
		fmt.Print("\n")
	}

	if x==10{
		fmt.Println("Stonks")
	}else{
		fmt.Println("No Stonks")
	}
	switch "Stonks"{
	case "S":
		fmt.Println("failed")
	case "Stonks":
		fmt.Println("failn't")
	default:
		fmt.Println("Printing default")
	}

	var numbers[3] string
	numbers[0]="one"
	numbers[1]="two"
	numbers[2]="three" 
	fmt.Println(numbers[1])
	fmt.Println(len(numbers))
	fmt.Println(numbers)

	mynewarray := [...] int {1,2,3,4,5}
	fmt.Println(mynewarray)

	var mynewarray2[] int = mynewarray[0:2]
	fmt.Print(mynewarray2)

	var var1[] int= append(mynewarray2,444)
	var var3[] int = append(mynewarray2,var1...)
	fmt.Println(var1)
	fmt.Println(var3)


	fmt.Println("\n##############Functions##################################\n")

	fmt.Println("this is function output : ",mynewfunc(1,2))

	newvar := calculation.do_add(1,2)
	fmt.Println(newvar)

}