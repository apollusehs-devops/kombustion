package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// SSMMaintenanceWindowTarget Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtarget.html
type SSMMaintenanceWindowTarget struct {
	Type       string                               `yaml:"Type"`
	Properties SSMMaintenanceWindowTargetProperties `yaml:"Properties"`
	Condition  interface{}                          `yaml:"Condition,omitempty"`
	Metadata   interface{}                          `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                          `yaml:"DependsOn,omitempty"`
}

// SSMMaintenanceWindowTarget Properties
type SSMMaintenanceWindowTargetProperties struct {
	Description      interface{} `yaml:"Description,omitempty"`
	Name             interface{} `yaml:"Name,omitempty"`
	OwnerInformation interface{} `yaml:"OwnerInformation,omitempty"`
	ResourceType     interface{} `yaml:"ResourceType"`
	WindowId         interface{} `yaml:"WindowId"`
	Targets          interface{} `yaml:"Targets"`
}

// NewSSMMaintenanceWindowTarget constructor creates a new SSMMaintenanceWindowTarget
func NewSSMMaintenanceWindowTarget(properties SSMMaintenanceWindowTargetProperties, deps ...interface{}) SSMMaintenanceWindowTarget {
	return SSMMaintenanceWindowTarget{
		Type:       "AWS::SSM::MaintenanceWindowTarget",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSSMMaintenanceWindowTarget parses SSMMaintenanceWindowTarget
func ParseSSMMaintenanceWindowTarget(
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
	var resource SSMMaintenanceWindowTarget
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
					"Fn::Sub": "${AWS::StackName}-SSMMaintenanceWindowTarget-" + name,
				},
			},
		},
	}

	return
}

// ParseSSMMaintenanceWindowTarget validator
func (resource SSMMaintenanceWindowTarget) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSSMMaintenanceWindowTargetProperties validator
func (resource SSMMaintenanceWindowTargetProperties) Validate() []error {
	errors := []error{}
	return errors
}