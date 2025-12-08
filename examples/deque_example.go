package main

import (
	"fmt"

	col "github.com/khajamoddin/collections/collections"
)

func dequeExample() {
	d := col.NewDeque[int]()
	d.PushBack(1)
	d.PushFront(0)
	v, _ := d.PopFront()
	fmt.Println(v)
}
