package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// MaintenanceWindowTaskTaskInvocationParameters Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html
type MaintenanceWindowTaskTaskInvocationParameters struct {
	MaintenanceWindowStepFunctionsParameters interface{} `yaml:"MaintenanceWindowStepFunctionsParameters,omitempty"`
	MaintenanceWindowRunCommandParameters    interface{} `yaml:"MaintenanceWindowRunCommandParameters,omitempty"`
	MaintenanceWindowLambdaParameters        interface{} `yaml:"MaintenanceWindowLambdaParameters,omitempty"`
	MaintenanceWindowAutomationParameters    interface{} `yaml:"MaintenanceWindowAutomationParameters,omitempty"`
}

// MaintenanceWindowTaskTaskInvocationParameters validation
func (resource MaintenanceWindowTaskTaskInvocationParameters) Validate() []error {
	errors := []error{}

	return errors
}
