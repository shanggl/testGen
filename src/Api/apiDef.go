package Api

import (
	"reflect"
	"Gen"
	"log"
)

type Foo struct{
	x string

}

func (f *Foo) SetX(v string ) {
	f.x=v
}

func (f * Foo) GetX() string {
	return f.x
}

func init(){
	log.Println("begin to register Type Of Foo")
	var f=Foo{}
	Gen.Register("Foo",reflect.TypeOf(f))

}
