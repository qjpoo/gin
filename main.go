package main

import (
	"fmt"
	uuid2 "github.com/satori/go.uuid"
	"strconv"
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


	t1, _ := strconv.ParseInt("3", 10, 64)
	fmt.Println(t1)
	t2, _ := strconv.Atoi("95")
	fmt.Println(t2)

}
