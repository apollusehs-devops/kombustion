package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// CloudFormationCustomResource Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cfn-customresource.html
type CloudFormationCustomResource struct {
	Type       string                                 `yaml:"Type"`
	Properties CloudFormationCustomResourceProperties `yaml:"Properties"`
	Condition  interface{}                            `yaml:"Condition,omitempty"`
	Metadata   interface{}                            `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                            `yaml:"DependsOn,omitempty"`
}

// CloudFormationCustomResource Properties
type CloudFormationCustomResourceProperties struct {
	ServiceToken interface{} `yaml:"ServiceToken"`
}

// NewCloudFormationCustomResource constructor creates a new CloudFormationCustomResource
func NewCloudFormationCustomResource(properties CloudFormationCustomResourceProperties, deps ...interface{}) CloudFormationCustomResource {
	return CloudFormationCustomResource{
		Type:       "AWS::CloudFormation::CustomResource",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCloudFormationCustomResource parses CloudFormationCustomResource
func ParseCloudFormationCustomResource(
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
	var resource CloudFormationCustomResource
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
					"Fn::Sub": "${AWS::StackName}-CloudFormationCustomResource-" + name,
				},
			},
		},
	}

	return
}

// ParseCloudFormationCustomResource validator
func (resource CloudFormationCustomResource) Validate() []error {
	return resource.Properties.Validate()
}

// ParseCloudFormationCustomResourceProperties validator
func (resource CloudFormationCustomResourceProperties) Validate() []error {
	errors := []error{}
	return errors
}
