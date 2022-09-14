package main

import (
	"fmt"
	"reflect"
)

// TODO| Разработать программу, которая в рантайме способна определить тип переменной:
// TODO| int, string, bool, channel из переменной типа interface{}.

func GetValueType(v any) { // можно так, но тогда для каждого типа канала нужен кейс
	switch v.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan")
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	GetValueType("Vitaly")
	GetValueType(24)
	GetValueType(true)
	GetValueType(make(chan int))

	var v any
	fmt.Println(reflect.TypeOf(v)) // А можно с помощью рефлексии

	v = make(chan int)
	fmt.Println(reflect.TypeOf(v).Kind().String())
	v = make(chan chan int)
	fmt.Println(reflect.TypeOf(v).Kind().String())
}
