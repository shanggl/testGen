package main

import (
	"Gen"
	"Api"
	"log"
	"reflect"
)

func main(){
	log.Println("begin")
	var f Api.Foo
	log.Println("foo in main " ,f)
	Gen.ShowTypes()
	p:=Gen.Gen("Foo")
	log.Println(reflect.TypeOf(p))
}

