package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// RulePredicate Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-rule-predicates.html
type RulePredicate struct {
	DataId  interface{} `yaml:"DataId"`
	Negated interface{} `yaml:"Negated"`
	Type    interface{} `yaml:"Type"`
}

// RulePredicate validation
func (resource RulePredicate) Validate() []error {
	errs := []error{}

	if resource.DataId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DataId'"))
	}
	if resource.Negated == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Negated'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}