package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// LoadBalancerPolicies Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-policy.html
type LoadBalancerPolicies struct {
	PolicyName        interface{} `yaml:"PolicyName"`
	PolicyType        interface{} `yaml:"PolicyType"`
	Attributes        interface{} `yaml:"Attributes"`
	InstancePorts     interface{} `yaml:"InstancePorts,omitempty"`
	LoadBalancerPorts interface{} `yaml:"LoadBalancerPorts,omitempty"`
}

// LoadBalancerPolicies validation
func (resource LoadBalancerPolicies) Validate() []error {
	errors := []error{}

	return errors
}
