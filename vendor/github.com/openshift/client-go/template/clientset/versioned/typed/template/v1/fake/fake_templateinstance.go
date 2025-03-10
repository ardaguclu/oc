// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/openshift/api/template/v1"
	templatev1 "github.com/openshift/client-go/template/applyconfigurations/template/v1"
	typedtemplatev1 "github.com/openshift/client-go/template/clientset/versioned/typed/template/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeTemplateInstances implements TemplateInstanceInterface
type fakeTemplateInstances struct {
	*gentype.FakeClientWithListAndApply[*v1.TemplateInstance, *v1.TemplateInstanceList, *templatev1.TemplateInstanceApplyConfiguration]
	Fake *FakeTemplateV1
}

func newFakeTemplateInstances(fake *FakeTemplateV1, namespace string) typedtemplatev1.TemplateInstanceInterface {
	return &fakeTemplateInstances{
		gentype.NewFakeClientWithListAndApply[*v1.TemplateInstance, *v1.TemplateInstanceList, *templatev1.TemplateInstanceApplyConfiguration](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("templateinstances"),
			v1.SchemeGroupVersion.WithKind("TemplateInstance"),
			func() *v1.TemplateInstance { return &v1.TemplateInstance{} },
			func() *v1.TemplateInstanceList { return &v1.TemplateInstanceList{} },
			func(dst, src *v1.TemplateInstanceList) { dst.ListMeta = src.ListMeta },
			func(list *v1.TemplateInstanceList) []*v1.TemplateInstance { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.TemplateInstanceList, items []*v1.TemplateInstance) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
