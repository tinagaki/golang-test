package main

import "fmt"

type Vertex struct {
	X int
	Y int
	S string
}
func changeVertex(v Vertex){
	v.X =1000
}
func changeVertex2(v *Vertex){
	v.X =1000
}

func main() {
	v:=Vertex{1,2,"test"}
	changeVertex(v)
	fmt.Println(v)
	v2:=&Vertex{1,2,"test"}
	changeVertex2(v2)
	fmt.Println(*v2)

	//v := Vertex{X: 1, Y: 2}
	//fmt.Println(v)
	//fmt.Println(v.X)
	//v.X = 100
	//fmt.Println(v.X)
	//v2 := Vertex{X: 1}
	//fmt.Println(v2)
	//v3 := Vertex{1, 2, "test"}
	//fmt.Println(v3)
	//
	//v4 := Vertex{}
	//fmt.Println(v4)
	//var v5 Vertex
	//fmt.Println(v5)
	//v6 := new(Vertex)
	//fmt.Printf("%T %v\n", v6, v6)
	//v7 := &Vertex{}
	//fmt.Printf("%T %v\n", v7, v7)
	//fmt.Println(v7.X)
	//
}

