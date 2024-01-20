package main

import (
         "fmt"
         "reflect"
)
type Person struct {
         Name string
         Age  int
}
func main() {
        person := Person{Name: "Alan", Age: 30}
        typ := reflect.TypeOf(person)
        zeroVal := reflect.Zero(typ)
        fmt.Println(zeroVal)
        
}