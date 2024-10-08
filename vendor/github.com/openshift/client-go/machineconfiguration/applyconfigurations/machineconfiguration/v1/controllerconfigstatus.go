// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ControllerConfigStatusApplyConfiguration represents a declarative configuration of the ControllerConfigStatus type for use
// with apply.
type ControllerConfigStatusApplyConfiguration struct {
	ObservedGeneration     *int64                                              `json:"observedGeneration,omitempty"`
	Conditions             []ControllerConfigStatusConditionApplyConfiguration `json:"conditions,omitempty"`
	ControllerCertificates []ControllerCertificateApplyConfiguration           `json:"controllerCertificates,omitempty"`
}

// ControllerConfigStatusApplyConfiguration constructs a declarative configuration of the ControllerConfigStatus type for use with
// apply.
func ControllerConfigStatus() *ControllerConfigStatusApplyConfiguration {
	return &ControllerConfigStatusApplyConfiguration{}
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *ControllerConfigStatusApplyConfiguration) WithObservedGeneration(value int64) *ControllerConfigStatusApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *ControllerConfigStatusApplyConfiguration) WithConditions(values ...*ControllerConfigStatusConditionApplyConfiguration) *ControllerConfigStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithControllerCertificates adds the given value to the ControllerCertificates field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ControllerCertificates field.
func (b *ControllerConfigStatusApplyConfiguration) WithControllerCertificates(values ...*ControllerCertificateApplyConfiguration) *ControllerConfigStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithControllerCertificates")
		}
		b.ControllerCertificates = append(b.ControllerCertificates, *values[i])
	}
	return b
}
