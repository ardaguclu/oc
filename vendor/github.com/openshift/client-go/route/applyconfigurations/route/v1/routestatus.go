// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// RouteStatusApplyConfiguration represents a declarative configuration of the RouteStatus type for use
// with apply.
type RouteStatusApplyConfiguration struct {
	Ingress []RouteIngressApplyConfiguration `json:"ingress,omitempty"`
}

// RouteStatusApplyConfiguration constructs a declarative configuration of the RouteStatus type for use with
// apply.
func RouteStatus() *RouteStatusApplyConfiguration {
	return &RouteStatusApplyConfiguration{}
}

// WithIngress adds the given value to the Ingress field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Ingress field.
func (b *RouteStatusApplyConfiguration) WithIngress(values ...*RouteIngressApplyConfiguration) *RouteStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithIngress")
		}
		b.Ingress = append(b.Ingress, *values[i])
	}
	return b
}
