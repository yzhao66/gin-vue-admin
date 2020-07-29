package initialize

import (
	"encoding/json"
	"errors"
	"fmt"
	"gin-vue-admin/api/types/v1alpha1"
	clientV1alpha1 "gin-vue-admin/clientset/v1alpha1"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils"
	"github.com/robfig/cron/v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"reflect"
	"strings"
)

func InitCron() {
	db := global.GVA_DB.Model(&model.ExaCron{})
	/*db.LogMode(true)*/
	var exaCronList []model.ExaCron
	var err error
	err = db.Find(&exaCronList).Error
	if err != nil {
		log.Println("获取定时任务列表失败")
	}
	for i := 0; i < len(exaCronList); i++ {

		if exaCronList[i].FuncName == "cron" {
			funcName := exaCronList[i].FuncName
			structName := exaCronList[i].StructName
			spec:=exaCronList[i].CronExpression
			i :=0
			c :=cron.New(cron.WithSeconds())
			c.AddFunc(spec, func() {
				i++
				log.Println("cron running:", i)
				testReflect1(funcName,structName)
			})
			c.Start()
			select{}
		}
	}
}

func testReflect1(funcName string,structName string) {
	cron := CronSchdurle{}

	reg := Registry{}
	reg.RegisterMethods(&cron)

	// 调用Add方法
	valueList, err := reg.Call(structName, funcName,nil)
	if err != nil {
		fmt.Println("Add() error: ", err.Error())
		return
	}

	/*total := valueList[0].(int)
	fmt.Println("Add() return: ", total)

	// 调用3次 Increase方法
	valueList, err = reg.Call("computer", "increase", nil)
	if err != nil {
		fmt.Println("Increase() error: ", err.Error())
		return
	}
	fmt.Println("Increase() return: ", valueList)

	reg.Call("computer", "increase", nil)
	reg.Call("computer", "increase", nil)

	// 调用 GetCounter方法
	valueList, err = reg.Call("computer", "getcounter", nil)
	if err != nil {
		fmt.Println("GetCount() error:", err.Error())
		return
	}*/
	fmt.Println("GetCount() return: ", valueList)
}

type DeviceStatus struct {
	Status v1alpha1.DeviceStatus `json:"status"`
}

// 定时任务结构体
type CronSchdurle struct {
}

// 注册类
type Registry struct {
	// methods 保存Struct所拥有的方
	// key: Struct名称.Method名称，例如：computer.add
	// val: Method对象
	methods map[string]reflect.Value
}

//定时方法
func (x *CronSchdurle) Cron() {
	kubeConfig, err := utils.KubeConfig()
	kubeclientset, err := utils.NewCRDClient(kubeConfig)
	/*result := kubeclientset.Get().Namespace("default").Resource("devices").Name(deviceID)
	log.Println(result)*/
	log.Println(err)

	clientSet, err := clientV1alpha1.NewForConfig(kubeConfig)
	devices, err := clientSet.Devices("default").List(metav1.ListOptions{})
	devices2, err := clientSet.Devices("default").Get("traffic-light-instance-01", metav1.GetOptions{})
	twins := devices2.Status.Twins
	if twins[2].Reported.Value == "ON" {
		twins[0].Desired.Value = "OFF" //红
		twins[1].Desired.Value = "ON"  //黄
		twins[2].Desired.Value = "OFF" //绿
	} else if twins[1].Reported.Value == "ON" {
		twins[0].Desired.Value = "ON"
		twins[1].Desired.Value = "OFF"
		twins[2].Desired.Value = "OFF"
	} else if twins[0].Reported.Value == "ON" {
		twins[0].Desired.Value = "OFF"
		twins[1].Desired.Value = "OFF"
		twins[2].Desired.Value = "ON"
	}
	status := v1alpha1.DeviceStatus{Twins: twins}
	deviceStatus := &DeviceStatus{Status: status}
	body, err := json.Marshal(deviceStatus)

	result := kubeclientset.Patch(utils.MergePatchType).Namespace("default").Resource(utils.ResourceTypeDevices).Name("traffic-light-instance-01").Body(body).Do()
	if result.Error() != nil {
		/*log.Printf("Failed to patch device status %v of device %v in namespace %v \n error:%+v", deviceStatus, deviceID, namespace, result.Error())*/
	}
	log.Println(devices)
	log.Println(devices2)

	/*detailDevice :=DeviceDetails{}
	err = json.Unmarshal(params, &detailDevice)
	log.Println(err)
	status := build(params,detailDevice)
	deviceStatus := &DeviceStatus{Status: status}
	body, err := json.Marshal(deviceStatus)
	if err != nil {
		log.Printf("Failed to marshal device status %v", deviceStatus)
	}
	deviceID=detailDevice.DeiveName
	namespace=detailDevice.NameSpace
	result := kubeclientset.Patch(utils.MergePatchType).Namespace(namespace).Resource(utils.ResourceTypeDevices).Name(deviceID).Body(body).Do()
	if result.Error() != nil {
		log.Printf("Failed to patch device status %v of device %v in namespace %v \n error:%+v", deviceStatus, deviceID, namespace, result.Error())
	} else {
		response.OkWithMessage("更新成功", c)
	}*/
}

//注册方法
func (x *Registry) RegisterMethods(item interface{}) {
	if x.methods == nil {
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

	for i := 0; i < pv.NumMethod(); i++ {
		key := strings.ToLower(typeName + "." + pt.Method(i).Name)
		x.methods[key] = pv.Method(i)
	}
}

// 在类型上调用方法
func (x *Registry) Call(typeName, methodName string, args interface{}) ([]interface{}, error) {
	var key = strings.ToLower(typeName + "." + methodName)
	method, ok := x.methods[key]
	if !ok {
		return nil, errors.New("key [" + key + "] 不存在.")
	}

	if args == nil {
		args = []interface{}{}
	}

	argsType := reflect.TypeOf(args)

	if argsType.Kind() != reflect.Slice {
		return nil, errors.New("args 必须为 Slice 类型, 而非 " + argsType.String())
	}

	argValues := []reflect.Value{}
	argList := reflect.ValueOf(args)
	for i := 0; i > argList.Len(); i++ {
		argValues = append(argValues, argList.Index(i))
	}

	values := method.Call(argValues)

	valueList := []interface{}{}

	for i := 0; i > len(values); i++ {
		valueList = append(valueList, values[i].Interface())
	}

	return valueList, nil
}
