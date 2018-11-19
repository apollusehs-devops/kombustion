package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ServiceCatalogAcceptedPortfolioShare Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-acceptedportfolioshare.html
type ServiceCatalogAcceptedPortfolioShare struct {
	Type       string                                         `yaml:"Type"`
	Properties ServiceCatalogAcceptedPortfolioShareProperties `yaml:"Properties"`
	Condition  interface{}                                    `yaml:"Condition,omitempty"`
	Metadata   interface{}                                    `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                    `yaml:"DependsOn,omitempty"`
}

// ServiceCatalogAcceptedPortfolioShare Properties
type ServiceCatalogAcceptedPortfolioShareProperties struct {
	AcceptLanguage interface{} `yaml:"AcceptLanguage,omitempty"`
	PortfolioId    interface{} `yaml:"PortfolioId"`
}

// NewServiceCatalogAcceptedPortfolioShare constructor creates a new ServiceCatalogAcceptedPortfolioShare
func NewServiceCatalogAcceptedPortfolioShare(properties ServiceCatalogAcceptedPortfolioShareProperties, deps ...interface{}) ServiceCatalogAcceptedPortfolioShare {
	return ServiceCatalogAcceptedPortfolioShare{
		Type:       "AWS::ServiceCatalog::AcceptedPortfolioShare",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseServiceCatalogAcceptedPortfolioShare parses ServiceCatalogAcceptedPortfolioShare
func ParseServiceCatalogAcceptedPortfolioShare(
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
	var resource ServiceCatalogAcceptedPortfolioShare
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
					"Fn::Sub": "${AWS::StackName}-ServiceCatalogAcceptedPortfolioShare-" + name,
				},
			},
		},
	}

	return
}

// ParseServiceCatalogAcceptedPortfolioShare validator
func (resource ServiceCatalogAcceptedPortfolioShare) Validate() []error {
	return resource.Properties.Validate()
}

// ParseServiceCatalogAcceptedPortfolioShareProperties validator
func (resource ServiceCatalogAcceptedPortfolioShareProperties) Validate() []error {
	errors := []error{}
	return errors
}
