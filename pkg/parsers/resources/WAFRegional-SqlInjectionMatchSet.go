package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// WAFRegionalSqlInjectionMatchSet Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sqlinjectionmatchset.html
type WAFRegionalSqlInjectionMatchSet struct {
	Type       string                                    `yaml:"Type"`
	Properties WAFRegionalSqlInjectionMatchSetProperties `yaml:"Properties"`
	Condition  interface{}                               `yaml:"Condition,omitempty"`
	Metadata   interface{}                               `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                               `yaml:"DependsOn,omitempty"`
}

// WAFRegionalSqlInjectionMatchSet Properties
type WAFRegionalSqlInjectionMatchSetProperties struct {
	Name                    interface{} `yaml:"Name"`
	SqlInjectionMatchTuples interface{} `yaml:"SqlInjectionMatchTuples,omitempty"`
}

// NewWAFRegionalSqlInjectionMatchSet constructor creates a new WAFRegionalSqlInjectionMatchSet
func NewWAFRegionalSqlInjectionMatchSet(properties WAFRegionalSqlInjectionMatchSetProperties, deps ...interface{}) WAFRegionalSqlInjectionMatchSet {
	return WAFRegionalSqlInjectionMatchSet{
		Type:       "AWS::WAFRegional::SqlInjectionMatchSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseWAFRegionalSqlInjectionMatchSet parses WAFRegionalSqlInjectionMatchSet
func ParseWAFRegionalSqlInjectionMatchSet(
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
	var resource WAFRegionalSqlInjectionMatchSet
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
					"Fn::Sub": "${AWS::StackName}-WAFRegionalSqlInjectionMatchSet-" + name,
				},
			},
		},
	}

	return
}

// ParseWAFRegionalSqlInjectionMatchSet validator
func (resource WAFRegionalSqlInjectionMatchSet) Validate() []error {
	return resource.Properties.Validate()
}

// ParseWAFRegionalSqlInjectionMatchSetProperties validator
func (resource WAFRegionalSqlInjectionMatchSetProperties) Validate() []error {
	errors := []error{}
	return errors
}
