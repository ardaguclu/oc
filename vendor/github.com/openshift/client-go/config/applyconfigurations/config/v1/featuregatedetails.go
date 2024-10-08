// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// FeatureGateDetailsApplyConfiguration represents a declarative configuration of the FeatureGateDetails type for use
// with apply.
type FeatureGateDetailsApplyConfiguration struct {
	Version  *string                                   `json:"version,omitempty"`
	Enabled  []FeatureGateAttributesApplyConfiguration `json:"enabled,omitempty"`
	Disabled []FeatureGateAttributesApplyConfiguration `json:"disabled,omitempty"`
}

// FeatureGateDetailsApplyConfiguration constructs a declarative configuration of the FeatureGateDetails type for use with
// apply.
func FeatureGateDetails() *FeatureGateDetailsApplyConfiguration {
	return &FeatureGateDetailsApplyConfiguration{}
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *FeatureGateDetailsApplyConfiguration) WithVersion(value string) *FeatureGateDetailsApplyConfiguration {
	b.Version = &value
	return b
}

// WithEnabled adds the given value to the Enabled field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Enabled field.
func (b *FeatureGateDetailsApplyConfiguration) WithEnabled(values ...*FeatureGateAttributesApplyConfiguration) *FeatureGateDetailsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithEnabled")
		}
		b.Enabled = append(b.Enabled, *values[i])
	}
	return b
}

// WithDisabled adds the given value to the Disabled field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Disabled field.
func (b *FeatureGateDetailsApplyConfiguration) WithDisabled(values ...*FeatureGateAttributesApplyConfiguration) *FeatureGateDetailsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithDisabled")
		}
		b.Disabled = append(b.Disabled, *values[i])
	}
	return b
}
