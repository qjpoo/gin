package main

import (
	"fmt"
	uuid2 "github.com/satori/go.uuid"
)

var b = 1

func f1()  {
	b:=2
	fmt.Println("f1 ...", b)
}

func f2()  {
	b:=3
	fmt.Println("f2 ...", b)


}
func main()  {
	//a = 2
	fmt.Println(b)

	f1()
	f2()

	uuid, e := uuid2.NewV4()
	fmt.Println(uuid.String(), e)

}
