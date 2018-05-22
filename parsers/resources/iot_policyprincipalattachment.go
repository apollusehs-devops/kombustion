package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type IoTPolicyPrincipalAttachment struct {
	Type       string                      `yaml:"Type"`
	Properties IoTPolicyPrincipalAttachmentProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type IoTPolicyPrincipalAttachmentProperties struct {
	PolicyName interface{} `yaml:"PolicyName"`
	Principal interface{} `yaml:"Principal"`
}

func NewIoTPolicyPrincipalAttachment(properties IoTPolicyPrincipalAttachmentProperties, deps ...interface{}) IoTPolicyPrincipalAttachment {
	return IoTPolicyPrincipalAttachment{
		Type:       "AWS::IoT::PolicyPrincipalAttachment",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseIoTPolicyPrincipalAttachment(name string, data string) (cf types.ValueMap, err error) {
	var resource IoTPolicyPrincipalAttachment
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: IoTPolicyPrincipalAttachment - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource IoTPolicyPrincipalAttachment) Validate() []error {
	return resource.Properties.Validate()
}

func (resource IoTPolicyPrincipalAttachmentProperties) Validate() []error {
	errs := []error{}
	if resource.PolicyName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PolicyName'"))
	}
	if resource.Principal == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Principal'"))
	}
	return errs
}
