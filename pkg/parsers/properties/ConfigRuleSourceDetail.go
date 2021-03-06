package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ConfigRuleSourceDetail Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source-sourcedetails.html
type ConfigRuleSourceDetail struct {
	EventSource               interface{} `yaml:"EventSource"`
	MaximumExecutionFrequency interface{} `yaml:"MaximumExecutionFrequency,omitempty"`
	MessageType               interface{} `yaml:"MessageType"`
}

// ConfigRuleSourceDetail validation
func (resource ConfigRuleSourceDetail) Validate() []error {
	errors := []error{}

	return errors
}
