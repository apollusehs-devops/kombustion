package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// FleetComputeCapacity Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-fleet-computecapacity.html
type FleetComputeCapacity struct {
	DesiredInstances interface{} `yaml:"DesiredInstances"`
}

// FleetComputeCapacity validation
func (resource FleetComputeCapacity) Validate() []error {
	errors := []error{}

	return errors
}
