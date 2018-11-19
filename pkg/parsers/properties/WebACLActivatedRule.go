package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// WebACLActivatedRule Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-webacl-rules.html
type WebACLActivatedRule struct {
	Priority interface{} `yaml:"Priority"`
	RuleId   interface{} `yaml:"RuleId"`
	Action   interface{} `yaml:"Action,omitempty"`
}

// WebACLActivatedRule validation
func (resource WebACLActivatedRule) Validate() []error {
	errors := []error{}

	return errors
}
