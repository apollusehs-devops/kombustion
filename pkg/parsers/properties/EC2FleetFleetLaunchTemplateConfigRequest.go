package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// EC2FleetFleetLaunchTemplateConfigRequest Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-fleetlaunchtemplateconfigrequest.html
type EC2FleetFleetLaunchTemplateConfigRequest struct {
	Overrides                   interface{} `yaml:"Overrides,omitempty"`
	LaunchTemplateSpecification interface{} `yaml:"LaunchTemplateSpecification,omitempty"`
}

// EC2FleetFleetLaunchTemplateConfigRequest validation
func (resource EC2FleetFleetLaunchTemplateConfigRequest) Validate() []error {
	errors := []error{}

	return errors
}