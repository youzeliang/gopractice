package main

import "fmt"

type Shape interface {
	Draw()
}

type Circle struct {
}

func (c *Circle) Draw() {
	fmt.Println("Drawing a circle...")
}

type Rectangle struct {
}

func (r *Rectangle) Draw() {
	fmt.Println("Drawing a rectangle...")
}

type ShapeFactory struct {
}

func (sf *ShapeFactory) CreateShape(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return &Circle{}
	case "rectangle":
		return &Rectangle{}
	default:
		return nil
	}
}

func main() {
	// 使用工厂类创建对象实例
	factory := &ShapeFactory{}
	circle := factory.CreateShape("circle")
	rectangle := factory.CreateShape("rectangle")
	// 调用对象方法
	circle.Draw()
	rectangle.Draw()
}
