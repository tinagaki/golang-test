package main

import "fmt"

func main() {
	//var n int = 100
	//var p *int = &n
	/*var p *int = new(int) //メモリ領域の確保
	fmt.Println(*p)
	*p++
	fmt.Println(*p)
*/
	/*
		var p2 *int
			fmt.Println(p2)　//nil
		*p2++
		fmt.Println(p2)
	*/
	s:= make([]int,0)
	fmt.Printf("%T\n",s)
	m:= make(map[string]int)
	fmt.Printf("%T\n",m)
	ch:= make(chan int)
	fmt.Printf("%T\n",ch)

	var p *int = new(int)
	fmt.Printf("%T\n",p)
	var st = new(struct{})
	fmt.Printf("%T\n",st)

}

