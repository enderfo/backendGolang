package main

import "fmt"

func main(){

for i := 0; i <= 4; i++{
	fmt.Println("Nilai i = ", i)
	if i==4 {
		for j := 0; j <= 10; j++ {
			if j == 5{
				unicode := "CAШAPBO"
				i:=0
				for index, runeValue := range unicode {
					fmt.Printf("character %#U starts at byte position %d\n", runeValue, i)
					_ = index
					i+=2
				}
				continue
			}
			fmt.Println("Nilai j = ", j)
		}
	}
}

}
