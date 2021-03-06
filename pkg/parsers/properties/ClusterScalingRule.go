package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ClusterScalingRule Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingrule.html
type ClusterScalingRule struct {
	Description interface{} `yaml:"Description,omitempty"`
	Name        interface{} `yaml:"Name"`
	Trigger     interface{} `yaml:"Trigger"`
	Action      interface{} `yaml:"Action"`
}

// ClusterScalingRule validation
func (resource ClusterScalingRule) Validate() []error {
	errors := []error{}

	return errors
}
