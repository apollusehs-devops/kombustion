package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ConfigRuleSource Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html
type ConfigRuleSource struct {
	Owner            interface{} `yaml:"Owner"`
	SourceIdentifier interface{} `yaml:"SourceIdentifier"`
	SourceDetails    interface{} `yaml:"SourceDetails,omitempty"`
}

// ConfigRuleSource validation
func (resource ConfigRuleSource) Validate() []error {
	errs := []error{}

	if resource.Owner == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Owner'"))
	}
	if resource.SourceIdentifier == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SourceIdentifier'"))
	}
	return errs
}