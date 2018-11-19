package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ClusterScalingAction Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingaction.html
type ClusterScalingAction struct {
	Market                           interface{} `yaml:"Market,omitempty"`
	SimpleScalingPolicyConfiguration interface{} `yaml:"SimpleScalingPolicyConfiguration"`
}

// ClusterScalingAction validation
func (resource ClusterScalingAction) Validate() []error {
	errors := []error{}

	return errors
}
