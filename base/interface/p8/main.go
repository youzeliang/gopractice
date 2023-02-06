package main

type Jedi interface {
	HasForce() bool
}

type Knight struct{}

func (k Knight) HasForce() bool {
	//TODO implement me
	panic("implement me")
}

var _ Jedi = (*Knight)(nil) // 利用编译器检查接口实现

func main() {}
