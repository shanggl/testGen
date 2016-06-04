package Gen
import (
	"reflect"
	"log"
)

var typeReg map[string] reflect.Type

func init(){
	typeReg=make(map[string] reflect.Type)
}


func ShowTypes(){
	for k,v := range typeReg{
		log.Println("_________________")
		log.Println(k,v)
	}

}

func Register(s string ,t reflect.Type){
	typeReg[s]=t
}
//gen a new Instance of s type
func Gen(s string ) interface{} {
	var ret interface{}
	t,ok:=typeReg[s]
	if !ok{
		return nil
	}
	ret=reflect.New(t).Interface()
	return ret
}

