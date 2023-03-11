package main

import "fmt"

func main(){
// 	fmt.Println("Hello World")

i := 21
fmt.Println(i)

fmt.Printf("%T \n",i)

fmt.Println(`%`)

var j bool = true
fmt.Printf("%t\n",j)

fmt.Println(!j)

fmt.Println("\u042f")

deci := 21
fmt.Printf("%d\n", deci)

octa := 025
fmt.Printf("%o\n",octa)

hexa := "f"
fmt.Printf("%s\n",hexa)

hexa2 := "F"
fmt.Printf("%s\n",hexa2)

const unicode2 = "Я"
for index, runeValue := range unicode2 {
	fmt.Printf("%U \n", runeValue)
	_ = index
}

var k float64 = 123.456
fmt.Printf("%f\n",k)

fmt.Printf("%E\n",k)
}