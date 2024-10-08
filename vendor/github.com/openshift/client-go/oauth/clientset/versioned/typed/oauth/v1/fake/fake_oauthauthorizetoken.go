// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/openshift/api/oauth/v1"
	oauthv1 "github.com/openshift/client-go/oauth/applyconfigurations/oauth/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOAuthAuthorizeTokens implements OAuthAuthorizeTokenInterface
type FakeOAuthAuthorizeTokens struct {
	Fake *FakeOauthV1
}

var oauthauthorizetokensResource = v1.SchemeGroupVersion.WithResource("oauthauthorizetokens")

var oauthauthorizetokensKind = v1.SchemeGroupVersion.WithKind("OAuthAuthorizeToken")

// Get takes name of the oAuthAuthorizeToken, and returns the corresponding oAuthAuthorizeToken object, and an error if there is any.
func (c *FakeOAuthAuthorizeTokens) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.OAuthAuthorizeToken, err error) {
	emptyResult := &v1.OAuthAuthorizeToken{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(oauthauthorizetokensResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.OAuthAuthorizeToken), err
}

// List takes label and field selectors, and returns the list of OAuthAuthorizeTokens that match those selectors.
func (c *FakeOAuthAuthorizeTokens) List(ctx context.Context, opts metav1.ListOptions) (result *v1.OAuthAuthorizeTokenList, err error) {
	emptyResult := &v1.OAuthAuthorizeTokenList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(oauthauthorizetokensResource, oauthauthorizetokensKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.OAuthAuthorizeTokenList{ListMeta: obj.(*v1.OAuthAuthorizeTokenList).ListMeta}
	for _, item := range obj.(*v1.OAuthAuthorizeTokenList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested oAuthAuthorizeTokens.
func (c *FakeOAuthAuthorizeTokens) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(oauthauthorizetokensResource, opts))
}

// Create takes the representation of a oAuthAuthorizeToken and creates it.  Returns the server's representation of the oAuthAuthorizeToken, and an error, if there is any.
func (c *FakeOAuthAuthorizeTokens) Create(ctx context.Context, oAuthAuthorizeToken *v1.OAuthAuthorizeToken, opts metav1.CreateOptions) (result *v1.OAuthAuthorizeToken, err error) {
	emptyResult := &v1.OAuthAuthorizeToken{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(oauthauthorizetokensResource, oAuthAuthorizeToken, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.OAuthAuthorizeToken), err
}

// Update takes the representation of a oAuthAuthorizeToken and updates it. Returns the server's representation of the oAuthAuthorizeToken, and an error, if there is any.
func (c *FakeOAuthAuthorizeTokens) Update(ctx context.Context, oAuthAuthorizeToken *v1.OAuthAuthorizeToken, opts metav1.UpdateOptions) (result *v1.OAuthAuthorizeToken, err error) {
	emptyResult := &v1.OAuthAuthorizeToken{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(oauthauthorizetokensResource, oAuthAuthorizeToken, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.OAuthAuthorizeToken), err
}

// Delete takes name of the oAuthAuthorizeToken and deletes it. Returns an error if one occurs.
func (c *FakeOAuthAuthorizeTokens) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(oauthauthorizetokensResource, name, opts), &v1.OAuthAuthorizeToken{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOAuthAuthorizeTokens) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(oauthauthorizetokensResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.OAuthAuthorizeTokenList{})
	return err
}

// Patch applies the patch and returns the patched oAuthAuthorizeToken.
func (c *FakeOAuthAuthorizeTokens) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.OAuthAuthorizeToken, err error) {
	emptyResult := &v1.OAuthAuthorizeToken{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(oauthauthorizetokensResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.OAuthAuthorizeToken), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied oAuthAuthorizeToken.
func (c *FakeOAuthAuthorizeTokens) Apply(ctx context.Context, oAuthAuthorizeToken *oauthv1.OAuthAuthorizeTokenApplyConfiguration, opts metav1.ApplyOptions) (result *v1.OAuthAuthorizeToken, err error) {
	if oAuthAuthorizeToken == nil {
		return nil, fmt.Errorf("oAuthAuthorizeToken provided to Apply must not be nil")
	}
	data, err := json.Marshal(oAuthAuthorizeToken)
	if err != nil {
		return nil, err
	}
	name := oAuthAuthorizeToken.Name
	if name == nil {
		return nil, fmt.Errorf("oAuthAuthorizeToken.Name must be provided to Apply")
	}
	emptyResult := &v1.OAuthAuthorizeToken{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(oauthauthorizetokensResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.OAuthAuthorizeToken), err
}
