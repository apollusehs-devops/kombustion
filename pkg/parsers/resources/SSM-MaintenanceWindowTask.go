package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// SSMMaintenanceWindowTask Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html
type SSMMaintenanceWindowTask struct {
	Type       string                             `yaml:"Type"`
	Properties SSMMaintenanceWindowTaskProperties `yaml:"Properties"`
	Condition  interface{}                        `yaml:"Condition,omitempty"`
	Metadata   interface{}                        `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                        `yaml:"DependsOn,omitempty"`
}

// SSMMaintenanceWindowTask Properties
type SSMMaintenanceWindowTaskProperties struct {
	Description              interface{} `yaml:"Description,omitempty"`
	MaxConcurrency           interface{} `yaml:"MaxConcurrency"`
	MaxErrors                interface{} `yaml:"MaxErrors"`
	Name                     interface{} `yaml:"Name,omitempty"`
	Priority                 interface{} `yaml:"Priority"`
	ServiceRoleArn           interface{} `yaml:"ServiceRoleArn"`
	TaskArn                  interface{} `yaml:"TaskArn"`
	TaskParameters           interface{} `yaml:"TaskParameters,omitempty"`
	TaskType                 interface{} `yaml:"TaskType"`
	WindowId                 interface{} `yaml:"WindowId,omitempty"`
	TaskInvocationParameters interface{} `yaml:"TaskInvocationParameters,omitempty"`
	LoggingInfo              interface{} `yaml:"LoggingInfo,omitempty"`
	Targets                  interface{} `yaml:"Targets"`
}

// NewSSMMaintenanceWindowTask constructor creates a new SSMMaintenanceWindowTask
func NewSSMMaintenanceWindowTask(properties SSMMaintenanceWindowTaskProperties, deps ...interface{}) SSMMaintenanceWindowTask {
	return SSMMaintenanceWindowTask{
		Type:       "AWS::SSM::MaintenanceWindowTask",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSSMMaintenanceWindowTask parses SSMMaintenanceWindowTask
func ParseSSMMaintenanceWindowTask(
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
	var resource SSMMaintenanceWindowTask
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
					"Fn::Sub": "${AWS::StackName}-SSMMaintenanceWindowTask-" + name,
				},
			},
		},
	}

	return
}

// ParseSSMMaintenanceWindowTask validator
func (resource SSMMaintenanceWindowTask) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSSMMaintenanceWindowTaskProperties validator
func (resource SSMMaintenanceWindowTaskProperties) Validate() []error {
	errors := []error{}
	return errors
}
