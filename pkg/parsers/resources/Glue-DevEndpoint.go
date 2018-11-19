package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// GlueDevEndpoint Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-devendpoint.html
type GlueDevEndpoint struct {
	Type       string                    `yaml:"Type"`
	Properties GlueDevEndpointProperties `yaml:"Properties"`
	Condition  interface{}               `yaml:"Condition,omitempty"`
	Metadata   interface{}               `yaml:"Metadata,omitempty"`
	DependsOn  interface{}               `yaml:"DependsOn,omitempty"`
}

// GlueDevEndpoint Properties
type GlueDevEndpointProperties struct {
	EndpointName          interface{} `yaml:"EndpointName,omitempty"`
	ExtraJarsS3Path       interface{} `yaml:"ExtraJarsS3Path,omitempty"`
	ExtraPythonLibsS3Path interface{} `yaml:"ExtraPythonLibsS3Path,omitempty"`
	NumberOfNodes         interface{} `yaml:"NumberOfNodes,omitempty"`
	PublicKey             interface{} `yaml:"PublicKey"`
	RoleArn               interface{} `yaml:"RoleArn"`
	SubnetId              interface{} `yaml:"SubnetId,omitempty"`
	SecurityGroupIds      interface{} `yaml:"SecurityGroupIds,omitempty"`
}

// NewGlueDevEndpoint constructor creates a new GlueDevEndpoint
func NewGlueDevEndpoint(properties GlueDevEndpointProperties, deps ...interface{}) GlueDevEndpoint {
	return GlueDevEndpoint{
		Type:       "AWS::Glue::DevEndpoint",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseGlueDevEndpoint parses GlueDevEndpoint
func ParseGlueDevEndpoint(
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
	var resource GlueDevEndpoint
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
					"Fn::Sub": "${AWS::StackName}-GlueDevEndpoint-" + name,
				},
			},
		},
	}

	return
}

// ParseGlueDevEndpoint validator
func (resource GlueDevEndpoint) Validate() []error {
	return resource.Properties.Validate()
}

// ParseGlueDevEndpointProperties validator
func (resource GlueDevEndpointProperties) Validate() []error {
	errors := []error{}
	return errors
}
