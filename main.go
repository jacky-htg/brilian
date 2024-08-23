package main

/*
import "fmt"

type Person struct {
	Name string
	Age  int
}

func (u Person) GetName() {
	println(u.Name)
}

type mystr string

func (u mystr) Println() {
	println(u)
}

func main() {
	var str mystr = "HAloa"
	str.Println()

	var joko Person = Person{
		Name: "joko",
		Age:  11,
	}
	fmt.Println(joko)
	joko.Name = "Joko"
	joko.Age = 12
	fmt.Println(joko)

	joko.GetName()

	cities := [...]string{"Jakarta", "Surabaya", "Bandung", "Jogjakarta", "Denpasar"}

	mySlice := cities[0:3]
	mySlice = append(mySlice, "Medan", "Padang", "Lampung", "Jakarta", "Surabaya", "Bandung", "Jogjakarta", "Denpasar", "Pontianak")
	fmt.Println(mySlice)
	fmt.Println(cities)
	println(len(mySlice), cap(mySlice))
	println(len(cities), cap(cities))

	var ages = make([]int, 3, 5)
	//ages = []int{41, 40, 50}
	fmt.Println(ages)
	println(len(ages), cap(ages))

	var hari map[string]int = map[string]int{"Senin": 1, "Selasa": 2, "Rabu": 3, "Kamis": 4, "Jumat": 3}
	fmt.Println(hari)

	/*for i, v := range "Jakarta" {
		println(i, v)
	}

	/*people := [...]Person{
		Person{
			Name: "Jet Lee",
			Age:  40,
		},
		Person{
			Name: "Jacky Chan",
			Age:  70,
		},
	}

	fmt.Printf("%+v", people)

	/*cities[0] = "Denpasar"
	cities[1] = "Surabaya"
	cities[2] = "Bandung"
	cities[3] = "Jogja"
	cities[4] = "Jakarta"

	fmt.Println(cities)
	fmt.Println(ages)
	println(cities[0])
}
*/
