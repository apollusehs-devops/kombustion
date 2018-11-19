package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// SESTemplate Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-template.html
type SESTemplate struct {
	Type       string                `yaml:"Type"`
	Properties SESTemplateProperties `yaml:"Properties"`
	Condition  interface{}           `yaml:"Condition,omitempty"`
	Metadata   interface{}           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}           `yaml:"DependsOn,omitempty"`
}

// SESTemplate Properties
type SESTemplateProperties struct {
	Template interface{} `yaml:"Template,omitempty"`
}

// NewSESTemplate constructor creates a new SESTemplate
func NewSESTemplate(properties SESTemplateProperties, deps ...interface{}) SESTemplate {
	return SESTemplate{
		Type:       "AWS::SES::Template",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSESTemplate parses SESTemplate
func ParseSESTemplate(
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
	var resource SESTemplate
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
					"Fn::Sub": "${AWS::StackName}-SESTemplate-" + name,
				},
			},
		},
	}

	return
}

// ParseSESTemplate validator
func (resource SESTemplate) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSESTemplateProperties validator
func (resource SESTemplateProperties) Validate() []error {
	errors := []error{}
	return errors
}
