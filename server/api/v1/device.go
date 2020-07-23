package v1

import (
	"encoding/json"
	"fmt"
	"gin-vue-admin/api/types/v1alpha1"
	clientV1alpha1 "gin-vue-admin/clientset/v1alpha1"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"log"
	"strconv"
	"time"
)

type JsonData struct {
	Data []Light
}

type Light struct {
	Color  string `json:"name"`
	Status string `json:"value"`
}
type DeviceStatus struct {
	Status v1alpha1.DeviceStatus `json:"status"`
}

var namespace = "default"
var deviceID = ""

func GetDeviceFile(c *gin.Context) {
	var result v1alpha1.Device
	var du model.Device
	_ = c.ShouldBindJSON(&du)
	kubeConfig, err := utils.KubeConfig()
	v1alpha1.AddToScheme(scheme.Scheme)
	clientSet, err := clientV1alpha1.NewForConfig(kubeConfig)
	if err != nil {
		panic(err)
	}
	devices, err := clientSet.Devices("default").List(metav1.ListOptions{})
	d := devices.Items
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(d); i++ {
		if(d[i].ObjectMeta.Name==du.DeviceName&&d[i].Spec.DeviceModelRef.Name==du.DeviceType && d[i].Spec.NodeSelector.NodeSelectorTerms[0].MatchExpressions[0].Values[0]==du.NodeName){
			result=d[i]
			log.Println(i)
			break
		}
	}
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), c)
	} else {
		response.OkWithData(result, c)
	}

}

func GetDeviceDetails(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	kubeConfig, err := utils.KubeConfig()
	v1alpha1.AddToScheme(scheme.Scheme)
	clientSet, err := clientV1alpha1.NewForConfig(kubeConfig)
	if err != nil {
		panic(err)
	}
	var deviceList []model.Device
	log.Println(deviceList)
	devices, err := clientSet.Devices("default").List(metav1.ListOptions{})
	d := devices.Items
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(d); i++ {
		var device model.Device
		device.DeviceType = d[i].Spec.DeviceModelRef.Name
		device.DeviceName = d[i].ObjectMeta.Name
		device.NodeName = d[i].Spec.NodeSelector.NodeSelectorTerms[0].MatchExpressions[0].Values[0]
		device.CreateTime = d[i].CreationTimestamp
		device.NameSpace=d[i].ObjectMeta.Namespace
		deviceList = append(deviceList, device)

	}
	//fmt.Printf("projects found: %+v\n", d)
	//c.JSON(http.StatusOK,d)

	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     deviceList,
			Total:    len(deviceList),
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
func UpdateDevice(c *gin.Context) {
	kubeConfig, err := utils.KubeConfig()
	kubeclientset, err := utils.NewCRDClient(kubeConfig)
	params, err := ioutil.ReadAll(c.Request.Body)
	detailDevice :=DeviceDetails{}
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
	}
}
func build(params []byte, device DeviceDetails) v1alpha1.DeviceStatus {
	/*var redValue string
	var greenValue string
	var yellowValue string*/
	metadata := map[string]string{"timestamp": strconv.FormatInt(time.Now().Unix()/1e6, 10),
		"type": "string",
	}
/*	jsonData := JsonData{}
	err := json.Unmarshal(params, &jsonData)
	log.Println(err)

	for i := 0; i < len(jsonData.Data); i++ {

		if jsonData.Data[i].Color == "red" {
			redValue = jsonData.Data[i].Status
		}
		if jsonData.Data[i].Color == "green" {
			greenValue = jsonData.Data[i].Status
		}
		if jsonData.Data[i].Color == "yellow" {
			yellowValue = jsonData.Data[i].Status
		}
	}*/

	twins := []v1alpha1.Twin{{PropertyName: "red", Desired: v1alpha1.TwinProperty{Value: device.Red, Metadata: metadata}, Reported: v1alpha1.TwinProperty{Value: device.Red, Metadata: metadata}}, {PropertyName: "green", Desired: v1alpha1.TwinProperty{Value: device.Green, Metadata: metadata}, Reported: v1alpha1.TwinProperty{Value: device.Green, Metadata: metadata}}, {PropertyName: "yellow", Desired: v1alpha1.TwinProperty{Value: device.Yellow, Metadata: metadata}, Reported: v1alpha1.TwinProperty{Value: device.Yellow, Metadata: metadata}}}
	devicestatus := v1alpha1.DeviceStatus{Twins: twins}
	return devicestatus
}

type DeviceDetails struct {
	Red  string `json:"red"`
	Yellow string `json:"yellow"`
	Green string `json:"green"`
	DeiveName string `json:"deiveName"`
	DeviceType string `json:"deviceType"`
	NodeName string `json:"nodeName"`
	NameSpace string `json:"nameSpace"`
}

func GetNodes(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	kubeConfig, err := utils.KubeConfig()
	Clientset, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		log.Printf("Failed to getnodes")
	}
	nodes, err := Clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	var nodeList []model.Node

	for i := 0; i < len(nodes.Items); i++ {
		var node model.Node
		node.NodeName = nodes.Items[i].ObjectMeta.Name
		nodeList = append(nodeList, node)
	}
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     nodeList,
			Total:    len(nodeList),
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}

}
func GetNode (c *gin.Context) {
	c = cron.New(cron.WithSeconds())
}