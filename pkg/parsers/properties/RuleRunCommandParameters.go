package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// RuleRunCommandParameters Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-runcommandparameters.html
type RuleRunCommandParameters struct {
	RunCommandTargets interface{} `yaml:"RunCommandTargets"`
}

// RuleRunCommandParameters validation
func (resource RuleRunCommandParameters) Validate() []error {
	errors := []error{}

	return errors
}
