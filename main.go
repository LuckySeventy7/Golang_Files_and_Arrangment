package main

import (
	"fmt"
	"os"
	"sort"
)

func ascendente(s []string){
	file1, err1 := os.Create("ascendente.txt")
	if err1 !=nil{
		fmt.Printf("error")
		return
	}

	defer file1.Close()

	sort.Strings(s)
	for _, j:= range s{
		file1.WriteString(j + "\n")
	}
}

func descendente(s []string){
	file2, err2 := os.Create("descendente.txt")
	if err2 !=nil{
		fmt.Printf("error")
		return
	}

	defer file2.Close()

	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	for _, j:= range s{
		file2.WriteString(j + "\n")
	}
}

func main(){
	
	var s []string
	var n int
	var cadena string
	fmt.Printf("Ingrese n:")
	fmt.Scan(&n)
	for i:=0; i < n; i++{
		fmt.Scan(&cadena)
		s = append(s, cadena)
	}
	ascendente(s)
	descendente(s)

	//for _,j:= range s{
	//	fmt.Print(j)
	//}

}