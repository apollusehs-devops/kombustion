package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// EC2FleetTargetCapacitySpecificationRequest Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-targetcapacityspecificationrequest.html
type EC2FleetTargetCapacitySpecificationRequest struct {
	DefaultTargetCapacityType interface{} `yaml:"DefaultTargetCapacityType,omitempty"`
	OnDemandTargetCapacity    interface{} `yaml:"OnDemandTargetCapacity,omitempty"`
	SpotTargetCapacity        interface{} `yaml:"SpotTargetCapacity,omitempty"`
	TotalTargetCapacity       interface{} `yaml:"TotalTargetCapacity"`
}

// EC2FleetTargetCapacitySpecificationRequest validation
func (resource EC2FleetTargetCapacitySpecificationRequest) Validate() []error {
	errors := []error{}

	return errors
}
