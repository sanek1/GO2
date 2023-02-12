package main

import (
	"fmt"
	"reflect"
)

type A struct {
	Value1  string
	Valuse2 string
	Note    string
}

func changeStruct(i interface{}, values map[string]interface{}) error {
	obj := reflect.Indirect(reflect.ValueOf(i))

	obj.FieldByName("Value1").SetString(values["v1"].(string))
	obj.FieldByName("Valuse2").SetString(values["v2"].(string))
	obj.FieldByName("Note").SetString(values["note"].(string))

	return fmt.Errorf("Done")
}

func main() {

	arr := map[string]interface{}{"v1": "change test 1", "v2": "change test 2", "note": "description"}
	a := A{Value1: "test1", Valuse2: "test2", Note: "Note"}

	fmt.Println(changeStruct(&a, arr))
	fmt.Println(a.Value1)
	fmt.Println(a.Valuse2)
	fmt.Println(a.Note)

}
