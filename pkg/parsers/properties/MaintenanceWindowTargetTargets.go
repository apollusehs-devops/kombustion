package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// MaintenanceWindowTargetTargets Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtarget-targets.html
type MaintenanceWindowTargetTargets struct {
	Key    interface{} `yaml:"Key"`
	Values interface{} `yaml:"Values,omitempty"`
}

// MaintenanceWindowTargetTargets validation
func (resource MaintenanceWindowTargetTargets) Validate() []error {
	errors := []error{}

	return errors
}