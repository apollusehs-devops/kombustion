package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// SageMakerModel Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-model.html
type SageMakerModel struct {
	Type       string                   `yaml:"Type"`
	Properties SageMakerModelProperties `yaml:"Properties"`
	Condition  interface{}              `yaml:"Condition,omitempty"`
	Metadata   interface{}              `yaml:"Metadata,omitempty"`
	DependsOn  interface{}              `yaml:"DependsOn,omitempty"`
}

// SageMakerModel Properties
type SageMakerModelProperties struct {
	ExecutionRoleArn interface{} `yaml:"ExecutionRoleArn"`
	ModelName        interface{} `yaml:"ModelName,omitempty"`
	VpcConfig        interface{} `yaml:"VpcConfig,omitempty"`
	Tags             interface{} `yaml:"Tags,omitempty"`
	PrimaryContainer interface{} `yaml:"PrimaryContainer"`
}

// NewSageMakerModel constructor creates a new SageMakerModel
func NewSageMakerModel(properties SageMakerModelProperties, deps ...interface{}) SageMakerModel {
	return SageMakerModel{
		Type:       "AWS::SageMaker::Model",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSageMakerModel parses SageMakerModel
func ParseSageMakerModel(
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
	var resource SageMakerModel
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
					"Fn::Sub": "${AWS::StackName}-SageMakerModel-" + name,
				},
			},
		},
	}

	return
}

// ParseSageMakerModel validator
func (resource SageMakerModel) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSageMakerModelProperties validator
func (resource SageMakerModelProperties) Validate() []error {
	errors := []error{}
	return errors
}
