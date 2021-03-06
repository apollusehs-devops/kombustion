package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// SageMakerEndpointConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html
type SageMakerEndpointConfig struct {
	Type       string                            `yaml:"Type"`
	Properties SageMakerEndpointConfigProperties `yaml:"Properties"`
	Condition  interface{}                       `yaml:"Condition,omitempty"`
	Metadata   interface{}                       `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                       `yaml:"DependsOn,omitempty"`
}

// SageMakerEndpointConfig Properties
type SageMakerEndpointConfigProperties struct {
	EndpointConfigName interface{} `yaml:"EndpointConfigName,omitempty"`
	KmsKeyId           interface{} `yaml:"KmsKeyId,omitempty"`
	ProductionVariants interface{} `yaml:"ProductionVariants"`
	Tags               interface{} `yaml:"Tags,omitempty"`
}

// NewSageMakerEndpointConfig constructor creates a new SageMakerEndpointConfig
func NewSageMakerEndpointConfig(properties SageMakerEndpointConfigProperties, deps ...interface{}) SageMakerEndpointConfig {
	return SageMakerEndpointConfig{
		Type:       "AWS::SageMaker::EndpointConfig",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSageMakerEndpointConfig parses SageMakerEndpointConfig
func ParseSageMakerEndpointConfig(
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
	var resource SageMakerEndpointConfig
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
					"Fn::Sub": "${AWS::StackName}-SageMakerEndpointConfig-" + name,
				},
			},
		},
	}

	return
}

// ParseSageMakerEndpointConfig validator
func (resource SageMakerEndpointConfig) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSageMakerEndpointConfigProperties validator
func (resource SageMakerEndpointConfigProperties) Validate() []error {
	errors := []error{}
	return errors
}
