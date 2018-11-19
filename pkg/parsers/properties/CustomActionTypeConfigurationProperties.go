package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// CustomActionTypeConfigurationProperties Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html
type CustomActionTypeConfigurationProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	Key         interface{} `yaml:"Key"`
	Name        interface{} `yaml:"Name"`
	Queryable   interface{} `yaml:"Queryable,omitempty"`
	Required    interface{} `yaml:"Required"`
	Secret      interface{} `yaml:"Secret"`
	Type        interface{} `yaml:"Type,omitempty"`
}

// CustomActionTypeConfigurationProperties validation
func (resource CustomActionTypeConfigurationProperties) Validate() []error {
	errors := []error{}

	return errors
}
