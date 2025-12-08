package main

import (
    "fmt"
    col "github.com/khajamoddin/collections/collections"
)

func main() {
    s := col.NewSet[int]()
    s.Add(1)
    s.Add(2)
    fmt.Println(s.Len())
}

