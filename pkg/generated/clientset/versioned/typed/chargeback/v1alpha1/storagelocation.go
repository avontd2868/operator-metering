package v1alpha1

import (
	v1alpha1 "github.com/coreos-inc/kube-chargeback/pkg/apis/chargeback/v1alpha1"
	scheme "github.com/coreos-inc/kube-chargeback/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// StorageLocationsGetter has a method to return a StorageLocationInterface.
// A group's client should implement this interface.
type StorageLocationsGetter interface {
	StorageLocations(namespace string) StorageLocationInterface
}

// StorageLocationInterface has methods to work with StorageLocation resources.
type StorageLocationInterface interface {
	Create(*v1alpha1.StorageLocation) (*v1alpha1.StorageLocation, error)
	Update(*v1alpha1.StorageLocation) (*v1alpha1.StorageLocation, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.StorageLocation, error)
	List(opts v1.ListOptions) (*v1alpha1.StorageLocationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageLocation, err error)
	StorageLocationExpansion
}

// storageLocations implements StorageLocationInterface
type storageLocations struct {
	client rest.Interface
	ns     string
}

// newStorageLocations returns a StorageLocations
func newStorageLocations(c *ChargebackV1alpha1Client, namespace string) *storageLocations {
	return &storageLocations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the storageLocation, and returns the corresponding storageLocation object, and an error if there is any.
func (c *storageLocations) Get(name string, options v1.GetOptions) (result *v1alpha1.StorageLocation, err error) {
	result = &v1alpha1.StorageLocation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("storagelocations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StorageLocations that match those selectors.
func (c *storageLocations) List(opts v1.ListOptions) (result *v1alpha1.StorageLocationList, err error) {
	result = &v1alpha1.StorageLocationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("storagelocations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested storageLocations.
func (c *storageLocations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("storagelocations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a storageLocation and creates it.  Returns the server's representation of the storageLocation, and an error, if there is any.
func (c *storageLocations) Create(storageLocation *v1alpha1.StorageLocation) (result *v1alpha1.StorageLocation, err error) {
	result = &v1alpha1.StorageLocation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("storagelocations").
		Body(storageLocation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a storageLocation and updates it. Returns the server's representation of the storageLocation, and an error, if there is any.
func (c *storageLocations) Update(storageLocation *v1alpha1.StorageLocation) (result *v1alpha1.StorageLocation, err error) {
	result = &v1alpha1.StorageLocation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("storagelocations").
		Name(storageLocation.Name).
		Body(storageLocation).
		Do().
		Into(result)
	return
}

// Delete takes name of the storageLocation and deletes it. Returns an error if one occurs.
func (c *storageLocations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("storagelocations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *storageLocations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("storagelocations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched storageLocation.
func (c *storageLocations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageLocation, err error) {
	result = &v1alpha1.StorageLocation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("storagelocations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
