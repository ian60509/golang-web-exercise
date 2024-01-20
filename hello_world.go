package main
import "fmt"
type Object struct {
	name string
}

type Rectengle struct {
	w, h int
	Object
}
type Circle struct {
	radius int
	Object
}

func (r Rectengle) area() int{
	return r.w * r.h
} 

func (c Circle) area() float64{
	r := (float64)(c.radius)
	return 3.14*r*r
}

func (c *Circle) setRadius(v int) {
	c.radius = v
}

func (o Object) sayHi() {
	fmt.Println("Hi my name is ",o.name)
}

func main(){
	var rr Rectengle = Rectengle{10,6,Object{"rr"}}
	cc := Circle{10,Object{"cc"}}
	cc.setRadius(1)
	fmt.Println(rr.area())
	fmt.Println(cc.area())
	cc.sayHi()
}