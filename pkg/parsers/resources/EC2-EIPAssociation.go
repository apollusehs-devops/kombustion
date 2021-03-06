package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// EC2EIPAssociation Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html
type EC2EIPAssociation struct {
	Type       string                      `yaml:"Type"`
	Properties EC2EIPAssociationProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

// EC2EIPAssociation Properties
type EC2EIPAssociationProperties struct {
	AllocationId       interface{} `yaml:"AllocationId,omitempty"`
	EIP                interface{} `yaml:"EIP,omitempty"`
	InstanceId         interface{} `yaml:"InstanceId,omitempty"`
	NetworkInterfaceId interface{} `yaml:"NetworkInterfaceId,omitempty"`
	PrivateIpAddress   interface{} `yaml:"PrivateIpAddress,omitempty"`
}

// NewEC2EIPAssociation constructor creates a new EC2EIPAssociation
func NewEC2EIPAssociation(properties EC2EIPAssociationProperties, deps ...interface{}) EC2EIPAssociation {
	return EC2EIPAssociation{
		Type:       "AWS::EC2::EIPAssociation",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2EIPAssociation parses EC2EIPAssociation
func ParseEC2EIPAssociation(
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
	var resource EC2EIPAssociation
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
					"Fn::Sub": "${AWS::StackName}-EC2EIPAssociation-" + name,
				},
			},
		},
	}

	return
}

// ParseEC2EIPAssociation validator
func (resource EC2EIPAssociation) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2EIPAssociationProperties validator
func (resource EC2EIPAssociationProperties) Validate() []error {
	errors := []error{}
	return errors
}
