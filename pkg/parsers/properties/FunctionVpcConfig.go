package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// FunctionVpcConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-vpcconfig.html
type FunctionVpcConfig struct {
	SecurityGroupIds interface{} `yaml:"SecurityGroupIds"`
	SubnetIds        interface{} `yaml:"SubnetIds"`
}

// FunctionVpcConfig validation
func (resource FunctionVpcConfig) Validate() []error {
	errors := []error{}

	return errors
}
