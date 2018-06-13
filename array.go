package main 

import (

   "fmt"
)


var s1 = [2]string{"hello","world"}
var s2 = [...]string{"hello","world","my first language on serve"}

func main() {

	fmt.Println("s1=",s1)

	fmt.Println("s2 = ",s2)


	c1 := make(chan [0]int)

	go func() {
		fmt.Println("C1")
		c1 <- [0]int{}
	}()
	<-c1 


	s := "hello,world"
	hello := s[:5]
	fmt.Println("hello=",hello)
	world := s[7:]
	fmt.Println("world=",world)


	fmt.Printf("%#v\n",[]byte("Hello,世界"))

	for i,r := range "Hello,世界" {
		fmt.Printf("%d\t%q\t%d\n",i,r,r)
	}

	fmt.Printf("%#v\n",str2bytes("My Name is fanlitao"))


	var a[]int

	a = append(a,1)
	a = append(a,2,3,4,5,6,7)
	a = append(a,[]int{8,9,10,11,12,13,14,15}...)
	for k,v:= range a {
		fmt.Printf("k=%d value=%d\n",k,v)
	}

	fmt.Println(".........................")
	a = append(a[:0],a[1:]...)

	for k,v := range a {
		fmt.Printf("k=%d value=%d\n",k,v)
	}

	fmt.Println("TrimSpace=",TrimSpace([]byte("Hello.  World.  Ni Hao")))

	for i := 0 ; i < 3;i++ {
		defer func() {println(i)}()
	}
}


func str2bytes(s string) []byte {
	p := make([]byte,len(s))
	for i := 0 ;i < len(s) ; i++ {
		c := s[i]
		p[i] = c
	}

	return p
}


func TrimSpace(s []byte) []byte {
	b := s[:0]
	for _,x := range s {
		if x != ' ' {
			b = append(b,x)
		}
	}
	return b
}