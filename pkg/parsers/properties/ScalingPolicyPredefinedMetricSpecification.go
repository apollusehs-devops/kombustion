package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ScalingPolicyPredefinedMetricSpecification Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predefinedmetricspecification.html
type ScalingPolicyPredefinedMetricSpecification struct {
	PredefinedMetricType interface{} `yaml:"PredefinedMetricType"`
	ResourceLabel        interface{} `yaml:"ResourceLabel,omitempty"`
}

// ScalingPolicyPredefinedMetricSpecification validation
func (resource ScalingPolicyPredefinedMetricSpecification) Validate() []error {
	errors := []error{}

	return errors
}
