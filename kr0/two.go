package main

import "fmt"

func main() {
	fmt.Printf("i am learning go\n")
	fmt.Printf("i am new\n")
	bow()
	fmt.Println("something more")
	{
		for k := 0; k < 10; k++ {
			if k%2 == 0 {
				fmt.Println(k)
			}
		}
		mow()
	}
}
func bow() {
	fmt.Printf("hello,i am boo\n")
	fmt.Printf("i am good. how about you?\n")
}
func mow() {
	fmt.Println("yes i got it and end")
}
