package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ModelVpcConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-vpcconfig.html
type ModelVpcConfig struct {
	SecurityGroupIds interface{} `yaml:"SecurityGroupIds"`
	Subnets          interface{} `yaml:"Subnets"`
}

// ModelVpcConfig validation
func (resource ModelVpcConfig) Validate() []error {
	errors := []error{}

	return errors
}
