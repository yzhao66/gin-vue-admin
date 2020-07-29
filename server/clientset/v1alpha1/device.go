package v1alpha1

import (
	"gin-vue-admin/api/types/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type DeviceInterface interface {
	List(opts metav1.ListOptions) (*v1alpha1.DeviceList, error)
	Get(name string, options metav1.GetOptions) (*v1alpha1.Device, error)
	Create(*v1alpha1.Device) (*v1alpha1.Device, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	// ...
}

type deviceClient struct {
	restClient rest.Interface
	ns         string
}

func (d deviceClient) List(opts metav1.ListOptions) (*v1alpha1.DeviceList, error) {
	result := v1alpha1.DeviceList{}
	err := d.restClient.
		Get().
		Namespace(d.ns).
		Resource("devices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (d deviceClient) Get(name string, options metav1.GetOptions) (*v1alpha1.Device, error) {
	result :=v1alpha1.Device{}
	err :=d.restClient.Get().Namespace(d.ns).Resource("devices").Name(name).Do().Into(&result)

	return &result,err
}

func (d deviceClient) Create(device *v1alpha1.Device) (*v1alpha1.Device, error) {
	panic("implement me")
}

func (d deviceClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	panic("implement me")
}


