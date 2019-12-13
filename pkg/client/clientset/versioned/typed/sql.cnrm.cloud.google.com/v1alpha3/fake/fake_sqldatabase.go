// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha3 "github.com/nais/naiserator/pkg/apis/sql.cnrm.cloud.google.com/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSqlDatabases implements SqlDatabaseInterface
type FakeSqlDatabases struct {
	Fake *FakeSqlV1alpha3
	ns   string
}

var sqldatabasesResource = schema.GroupVersionResource{Group: "sql.cnrm.cloud.google.com", Version: "v1alpha3", Resource: "sqldatabases"}

var sqldatabasesKind = schema.GroupVersionKind{Group: "sql.cnrm.cloud.google.com", Version: "v1alpha3", Kind: "SqlDatabase"}

// Get takes name of the sqlDatabase, and returns the corresponding sqlDatabase object, and an error if there is any.
func (c *FakeSqlDatabases) Get(name string, options v1.GetOptions) (result *v1alpha3.SqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sqldatabasesResource, c.ns, name), &v1alpha3.SqlDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.SqlDatabase), err
}

// List takes label and field selectors, and returns the list of SqlDatabases that match those selectors.
func (c *FakeSqlDatabases) List(opts v1.ListOptions) (result *v1alpha3.SqlDatabaseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sqldatabasesResource, sqldatabasesKind, c.ns, opts), &v1alpha3.SqlDatabaseList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha3.SqlDatabaseList{ListMeta: obj.(*v1alpha3.SqlDatabaseList).ListMeta}
	for _, item := range obj.(*v1alpha3.SqlDatabaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sqlDatabases.
func (c *FakeSqlDatabases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sqldatabasesResource, c.ns, opts))

}

// Create takes the representation of a sqlDatabase and creates it.  Returns the server's representation of the sqlDatabase, and an error, if there is any.
func (c *FakeSqlDatabases) Create(sqlDatabase *v1alpha3.SqlDatabase) (result *v1alpha3.SqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sqldatabasesResource, c.ns, sqlDatabase), &v1alpha3.SqlDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.SqlDatabase), err
}

// Update takes the representation of a sqlDatabase and updates it. Returns the server's representation of the sqlDatabase, and an error, if there is any.
func (c *FakeSqlDatabases) Update(sqlDatabase *v1alpha3.SqlDatabase) (result *v1alpha3.SqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sqldatabasesResource, c.ns, sqlDatabase), &v1alpha3.SqlDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.SqlDatabase), err
}

// Delete takes name of the sqlDatabase and deletes it. Returns an error if one occurs.
func (c *FakeSqlDatabases) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sqldatabasesResource, c.ns, name), &v1alpha3.SqlDatabase{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSqlDatabases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sqldatabasesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha3.SqlDatabaseList{})
	return err
}

// Patch applies the patch and returns the patched sqlDatabase.
func (c *FakeSqlDatabases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha3.SqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sqldatabasesResource, c.ns, name, pt, data, subresources...), &v1alpha3.SqlDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.SqlDatabase), err
}
