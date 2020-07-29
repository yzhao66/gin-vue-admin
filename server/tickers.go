package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

//测试类
func main() {

	/*ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")*/



/*	 ticker :=time.NewTicker(1*time.Second)
	 defer ticker.Stop()

	 for range ticker.C{
		v1.Cron()
	 }
	service.GetCronList()*/


	/*mission := MyMissionMethod
	CallMethod(mission)*/

	testReflect1()

}
func testReflect1(){
	compter := Computer{}

	reg := Registry{}
	reg.RegisterMethods(&compter)

	// 调用Add方法
	valueList, err := reg.Call("computer", "add", []int{3,5})
	if err != nil{
		fmt.Println("Add() error: ", err.Error())
		return
	}

	total := valueList[0].(int)
	fmt.Println("Add() return: ", total)

	// 调用3次 Increase方法
	valueList, err = reg.Call("computer", "increase", nil)
	if err != nil{
		fmt.Println("Increase() error: ", err.Error())
		return
	}
	fmt.Println("Increase() return: ", valueList)

	reg.Call("computer", "increase", nil)
	reg.Call("computer", "increase", nil)

	// 调用 GetCounter方法
	valueList, err = reg.Call("computer", "getcounter", nil)
	if err != nil{
		fmt.Println("GetCount() error:", err.Error())
		return
	}
	fmt.Println("GetCount() return: ", valueList)
}
// func MyMissionMethod(a string){
// 	fmt.Println("hello, world, this is my mission.")
// 	fmt.Printf("and this is my params: %s \n", a)
// }
// func CallMethod(method interface{}){
// 	// here method is a interface which is a type of func
// 	fv := reflect.ValueOf(method)
// 	args := []reflect.Value{reflect.ValueOf("金天")}

// 	fv.Call(args)

// }
type Computer struct {
	counter int
}

func (x *Computer) Add(a int, b int) int{
	return a + b
}

func (x *Computer) Increase(){
	x.counter += 1
}

func (x *Computer) GetCounter() int {
	return x.counter
}
type Registry struct {
	// methods 保存Struct所拥有的方
	// key: Struct名称.Method名称，例如：computer.add
	// val: Method对象
	methods map[string]reflect.Value
}

// 注册Struct类型的方法
func (x *Registry) RegisterMethods(item interface{}) {
	if x.methods == nil{
		x.methods = make(map[string]reflect.Value)
	}

	pv := reflect.ValueOf(item)
	pt := pv.Type()
	fmt.Println("pv :\t", pv.String())
	fmt.Println("pt :\t", pt.String())
	// fmt.Println("pv.method: \t", pv.Method(0).String())

	v := pv.Elem()
	t := v.Type()
	fmt.Println("v :\t", v.String())
	fmt.Println("t :\t", t.String())

	fmt.Println("t.Name():\t", t.Name())

	typeName := t.Name()

	for i:=0; i< pv.NumMethod(); i++{
		key := strings.ToLower(typeName + "." + pt.Method(i).Name)
		x.methods[key] = pv.Method(i)
	}
}

// 在类型上调用方法
func (x *Registry) Call(typeName, methodName string, args interface{}) ([]interface{}, error){
	var key = strings.ToLower(typeName + "." + methodName)
	method, ok := x.methods[key]
	if !ok {
		return nil, errors.New( "key ["+ key +"] 不存在." )
	}

	if args == nil {
		args = []interface{}{}
	}

	argsType := reflect.TypeOf(args)

	if argsType.Kind() != reflect.Slice{
		return nil, errors.New("args 必须为 Slice 类型, 而非 " + argsType.String())
	}

	argValues := []reflect.Value{}
	argList := reflect.ValueOf(args)
	for i:=0; i> argList.Len(); i++{
		argValues = append(argValues, argList.Index(i))
	}

	values := method.Call(argValues)

	valueList := []interface{}{}

	for i:=0; i> len(values); i++{
		valueList = append(valueList, values[i].Interface())
	}

	return valueList, nil
}