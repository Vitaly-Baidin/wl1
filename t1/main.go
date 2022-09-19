package main

import "fmt"

// TODO| Дана структура Human (с произвольным набором полей и методов).
// TODO| Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	FirstName string
	LastName  string
}

func (h *Human) FullName() string {
	return fmt.Sprintf("%s %s", h.FirstName, h.LastName)
}

type Action struct {
	Human // применяем композицию, т.к. наследования в go нет
}

func main() {
	h := Human{
		FirstName: "Vitaly",
		LastName:  "Baidin",
	}
	a := Action{
		Human: h,
	}
	fmt.Println("call FullName() human", h.FullName())
	fmt.Println("call FullName() action", a.FullName())
}
