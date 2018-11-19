package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// GuardDutyThreatIntelSet Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-threatintelset.html
type GuardDutyThreatIntelSet struct {
	Type       string                            `yaml:"Type"`
	Properties GuardDutyThreatIntelSetProperties `yaml:"Properties"`
	Condition  interface{}                       `yaml:"Condition,omitempty"`
	Metadata   interface{}                       `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                       `yaml:"DependsOn,omitempty"`
}

// GuardDutyThreatIntelSet Properties
type GuardDutyThreatIntelSetProperties struct {
	Activate   interface{} `yaml:"Activate"`
	DetectorId interface{} `yaml:"DetectorId"`
	Format     interface{} `yaml:"Format"`
	Location   interface{} `yaml:"Location"`
	Name       interface{} `yaml:"Name,omitempty"`
}

// NewGuardDutyThreatIntelSet constructor creates a new GuardDutyThreatIntelSet
func NewGuardDutyThreatIntelSet(properties GuardDutyThreatIntelSetProperties, deps ...interface{}) GuardDutyThreatIntelSet {
	return GuardDutyThreatIntelSet{
		Type:       "AWS::GuardDuty::ThreatIntelSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseGuardDutyThreatIntelSet parses GuardDutyThreatIntelSet
func ParseGuardDutyThreatIntelSet(
	name string,
	data string,
) (
	source string,
	conditions types.TemplateObject,
	metadata types.TemplateObject,
	mappings types.TemplateObject,
	outputs types.TemplateObject,
	parameters types.TemplateObject,
	resources types.TemplateObject,
	transform types.TemplateObject,
	errors []error,
) {
	source = "kombustion-core-resources"

	// Resources
	var resource GuardDutyThreatIntelSet
	err := yaml.Unmarshal([]byte(data), &resource)

	if err != nil {
		errors = append(errors, err)
		return
	}

	if validateErrs := resource.Properties.Validate(); len(errors) > 0 {
		errors = append(errors, validateErrs...)
		return
	}

	resources = types.TemplateObject{name: resource}

	// Outputs

	outputs = types.TemplateObject{
		name: types.TemplateObject{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-GuardDutyThreatIntelSet-" + name,
				},
			},
		},
	}

	return
}

// ParseGuardDutyThreatIntelSet validator
func (resource GuardDutyThreatIntelSet) Validate() []error {
	return resource.Properties.Validate()
}

// ParseGuardDutyThreatIntelSetProperties validator
func (resource GuardDutyThreatIntelSetProperties) Validate() []error {
	errors := []error{}
	return errors
}
