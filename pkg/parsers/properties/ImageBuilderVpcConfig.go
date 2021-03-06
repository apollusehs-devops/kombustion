package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ImageBuilderVpcConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-imagebuilder-vpcconfig.html
type ImageBuilderVpcConfig struct {
	SecurityGroupIds interface{} `yaml:"SecurityGroupIds,omitempty"`
	SubnetIds        interface{} `yaml:"SubnetIds,omitempty"`
}

// ImageBuilderVpcConfig validation
func (resource ImageBuilderVpcConfig) Validate() []error {
	errors := []error{}

	return errors
}
