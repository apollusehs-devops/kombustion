package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ScalingPolicyCustomizedMetricSpecification Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html
type ScalingPolicyCustomizedMetricSpecification struct {
	MetricName interface{} `yaml:"MetricName"`
	Namespace  interface{} `yaml:"Namespace"`
	Statistic  interface{} `yaml:"Statistic"`
	Unit       interface{} `yaml:"Unit,omitempty"`
	Dimensions interface{} `yaml:"Dimensions,omitempty"`
}

// ScalingPolicyCustomizedMetricSpecification validation
func (resource ScalingPolicyCustomizedMetricSpecification) Validate() []error {
	errors := []error{}

	return errors
}
