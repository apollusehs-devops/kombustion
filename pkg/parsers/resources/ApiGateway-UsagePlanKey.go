package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ApiGatewayUsagePlanKey Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplankey.html
type ApiGatewayUsagePlanKey struct {
	Type       string                           `yaml:"Type"`
	Properties ApiGatewayUsagePlanKeyProperties `yaml:"Properties"`
	Condition  interface{}                      `yaml:"Condition,omitempty"`
	Metadata   interface{}                      `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                      `yaml:"DependsOn,omitempty"`
}

// ApiGatewayUsagePlanKey Properties
type ApiGatewayUsagePlanKeyProperties struct {
	KeyId       interface{} `yaml:"KeyId"`
	KeyType     interface{} `yaml:"KeyType"`
	UsagePlanId interface{} `yaml:"UsagePlanId"`
}

// NewApiGatewayUsagePlanKey constructor creates a new ApiGatewayUsagePlanKey
func NewApiGatewayUsagePlanKey(properties ApiGatewayUsagePlanKeyProperties, deps ...interface{}) ApiGatewayUsagePlanKey {
	return ApiGatewayUsagePlanKey{
		Type:       "AWS::ApiGateway::UsagePlanKey",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApiGatewayUsagePlanKey parses ApiGatewayUsagePlanKey
func ParseApiGatewayUsagePlanKey(
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
	var resource ApiGatewayUsagePlanKey
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
					"Fn::Sub": "${AWS::StackName}-ApiGatewayUsagePlanKey-" + name,
				},
			},
		},
	}

	return
}

// ParseApiGatewayUsagePlanKey validator
func (resource ApiGatewayUsagePlanKey) Validate() []error {
	return resource.Properties.Validate()
}

// ParseApiGatewayUsagePlanKeyProperties validator
func (resource ApiGatewayUsagePlanKeyProperties) Validate() []error {
	errors := []error{}
	return errors
}
