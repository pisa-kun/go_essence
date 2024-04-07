package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	var n int = 1
	fmt.Println(n)
	n = 2
	fmt.Println(n)
	// 型が違う値は代入不可
	// n = "3"

	// var s = "string"と同義
	s := "string"
	fmt.Println(s)

	// iota
	type Hokura int

	const (
		Kaho Hokura = iota
		Chyoko
		Juri
		Rinze
		Natsuha
	)
	var member Hokura = Natsuha
	fmt.Println(member)

	// slice
	a := []int{1, 2, 3}
	fmt.Println(a)
	a = append(a, 4, 5, 6)
	fmt.Println(a)

	// map
	m := map[string]int{
		"kaho":    1,
		"chyoko":  2,
		"juri":    3,
		"rinze":   4,
		"natsuha": 5,
	}
	fmt.Println(m)
	fmt.Println(m["natsuha"])

	for k, v := range m {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}

	// 存在しないkey
	// 0を出力
	fmt.Println(m["mamimi"])
	v, ok := m["natsuha"]
	if ok {
		fmt.Println(v)
	}

	var suzumoto Member
	suzumoto.Name = "suzumoto"
	suzumoto.Age = 22
	fmt.Println(suzumoto)
	showName(&suzumoto)

	// interface
	var speaker Speaker = &Dog{}
	DoSpeak(speaker)
	speaker = &Cat{}
	DoSpeak(speaker)
}

type Member struct {
	Name string
	Age  int
}

func showName(member *Member) {
	fmt.Println(member.Name)
}

type Speaker interface {
	Speak() error
}

type Dog struct{}
type Cat struct{}

func (d *Dog) Speak() error {
	fmt.Println("Bow")
	return nil
}
func (c *Cat) Speak() error {
	fmt.Println("Meow")
	return nil
}
func DoSpeak(s Speaker) error {
	return s.Speak()
}
