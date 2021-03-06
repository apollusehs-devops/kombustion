package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ServiceDeploymentConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentconfiguration.html
type ServiceDeploymentConfiguration struct {
	MaximumPercent        interface{} `yaml:"MaximumPercent,omitempty"`
	MinimumHealthyPercent interface{} `yaml:"MinimumHealthyPercent,omitempty"`
}

// ServiceDeploymentConfiguration validation
func (resource ServiceDeploymentConfiguration) Validate() []error {
	errors := []error{}

	return errors
}
