package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jacky-htg/brilian/lib"
)

func MyMethod() {
	defer func() {
		if err := recover(); err != nil {
			println("from revcover", err)
		}
	}()

	if true {
		panic("PANIC")
	}
}

type operation func(int, int) int

type Person struct {
	Name string
	Age  int
}

type anything interface {
}

type myI interface {
	Operation() int
	Hello()
	Print()
}

type MyDiff int
type MyAdd int

func (u MyDiff) Operation() int {
	return 10
}

func (u MyDiff) Hello() {
	println("hello from mydiff")
}

func (u MyDiff) Print() {
	println("print from mydiff")
}

func (u *Person) PrintName() {
	println(u.Name)
}

func (u *Person) Operation() int {
	return 20
}

func (u *Person) Hello() {
	println("Hello from Person")
}

func (u *Person) Print() {
	println("print  from Person")
}

func someErr() error {
	return errors.New("err from someErr")
}

func errFmt() error {
	return fmt.Errorf("error from %v", []int{1, 4, 5})
}

func main() {
	logErr := log.New(os.Stdout, "ERROR : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	info := log.New(os.Stdout, "INFO : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	MyMethod()

	info.Println("Some Info")

	err := someErr()
	if err != nil {
		logErr.Println(err.Error())
		return
	}

	err = errFmt()
	if err != nil {
		println(err.Error())
		return
	}

	//println(app.Ctx("Halo halo bandung"))
	println(lib.MyStr + lib.Pagi())
	var salam lib.Salam
	salam.Message = "Good"
	salam.Type = "Night"
	println(salam.Sapa())
	//println(db.IsConn())
	/*myarr := []any{true, "satu", 2, Person{}, MyDiff(4)}
	fmt.Println(myarr)
	println(myarr[0].(bool))

	var i myI = MyDiff(10)
	println(i.Operation())
	i.Hello()

	var i2 myI = &Person{}
	println(i2.Operation())
	i2.Hello()

	var d MyDiff
	println(d.Operation())

	var budi Person = Person{Name: "Budi", Age: 34}
	//fmt.Printf("%T", budi)
	//budi.Name = "Budi"
	//budi.Age = 20
	budi.PrintName()
	println(budi.Name, budi.Age)
	fmt.Printf("%+v", budi)

	var people []Person = []Person{
		{Name: "Budi", Age: 4},
		{Name: "Wati", Age: 12},
	}

	people = append(people, Person{Name: "Dika", Age: 34})
	fmt.Println(people)

	/*var a int
	a = 10
	println(a)
	a = 20
	println(a)

	b := &a
	println(b)
	println(*b)
	*b = 30
	println(*b, a)

	var mystr string
	getStr(&mystr)
	println(mystr)

	/*println("begin")
	defer func() {
		println("halo")
	}()

	for i := 1; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}

	println("end")

	/*result := add(2, 5)
	println(result)

	addFunc := func(a, b int) int {
		return a + b
	}

	var diffFunc operation = func(a, b int) int {
		return a - b
	}

	fmt.Printf("%T", addFunc)
	println("")
	fmt.Printf("%T", diffFunc)
	println("")

	println(addFunc(5, 8))
	println(addFunc(4, 8))
	println(diffFunc)

	println(myOperation(add, 2, 4))
	println(myOperation(addFunc, 2, 4))
	println(myOperation(diffFunc, 2, 4))

	mySlice := []int{12, 3, 5, 6}
	println(sum(11, 2, 5))
	println("halo", "apa kabar", "aloha")
	println(sum(23, mySlice...))*/
}

func myOperation(f operation, a, b int) int {
	return f(a, b)
}

func add(a, b int) int {
	return a + b
}

func sum(a int, args ...int) int {
	var total int
	println(a)
	for _, v := range args {
		total += v
	}
	return total
}

func getStr(str *string) {
	*str = "halo"
}
