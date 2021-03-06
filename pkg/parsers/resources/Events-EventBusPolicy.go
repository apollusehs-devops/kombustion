package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// EventsEventBusPolicy Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbuspolicy.html
type EventsEventBusPolicy struct {
	Type       string                         `yaml:"Type"`
	Properties EventsEventBusPolicyProperties `yaml:"Properties"`
	Condition  interface{}                    `yaml:"Condition,omitempty"`
	Metadata   interface{}                    `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                    `yaml:"DependsOn,omitempty"`
}

// EventsEventBusPolicy Properties
type EventsEventBusPolicyProperties struct {
	Action      interface{} `yaml:"Action"`
	Principal   interface{} `yaml:"Principal"`
	StatementId interface{} `yaml:"StatementId"`
	Condition   interface{} `yaml:"Condition,omitempty"`
}

// NewEventsEventBusPolicy constructor creates a new EventsEventBusPolicy
func NewEventsEventBusPolicy(properties EventsEventBusPolicyProperties, deps ...interface{}) EventsEventBusPolicy {
	return EventsEventBusPolicy{
		Type:       "AWS::Events::EventBusPolicy",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEventsEventBusPolicy parses EventsEventBusPolicy
func ParseEventsEventBusPolicy(
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
	var resource EventsEventBusPolicy
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
					"Fn::Sub": "${AWS::StackName}-EventsEventBusPolicy-" + name,
				},
			},
		},
	}

	return
}

// ParseEventsEventBusPolicy validator
func (resource EventsEventBusPolicy) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEventsEventBusPolicyProperties validator
func (resource EventsEventBusPolicyProperties) Validate() []error {
	errors := []error{}
	return errors
}
