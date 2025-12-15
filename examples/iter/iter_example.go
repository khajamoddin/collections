package main

import (
	"fmt"

	col "github.com/khajamoddin/collections/collections"
)

func main() {
	// Deque Iterator
	d := col.NewDeque[int]()
	d.PushBack(1)
	d.PushBack(2)
	fmt.Print("Deque: ")
	for v := range d.All() {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	// Set Iterator
	s := col.NewSet[string]()
	s.Add("a")
	s.Add("b")
	fmt.Print("Set: ")
	for v := range s.All() {
		fmt.Printf("%s ", v)
	}
	fmt.Println()

	// OrderedMap Iterator
	om := col.NewOrderedMap[string, int]()
	om.Set("one", 1)
	om.Set("two", 2)
	fmt.Print("OrderedMap: ")
	for k, v := range om.All() {
		fmt.Printf("%s:%d ", k, v)
	}
	fmt.Println()

	// MultiMap Iterator
	mm := col.NewMultiMap[string, int]()
	mm.Add("nums", 10)
	mm.Add("nums", 20)
	fmt.Print("MultiMap: ")
	for k, v := range mm.All() {
		fmt.Printf("%s:%d ", k, v)
	}
	fmt.Println()
}
