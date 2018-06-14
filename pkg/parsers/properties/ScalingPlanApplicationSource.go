package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ScalingPlanApplicationSource Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-applicationsource.html
type ScalingPlanApplicationSource struct {
	CloudFormationStackARN interface{} `yaml:"CloudFormationStackARN,omitempty"`
	TagFilters             interface{} `yaml:"TagFilters,omitempty"`
}

// ScalingPlanApplicationSource validation
func (resource ScalingPlanApplicationSource) Validate() []error {
	errs := []error{}

	return errs
}