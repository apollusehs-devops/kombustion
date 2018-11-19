package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ScalingPolicyTargetTrackingConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingconfiguration.html
type ScalingPolicyTargetTrackingConfiguration struct {
	DisableScaleIn                interface{} `yaml:"DisableScaleIn,omitempty"`
	TargetValue                   interface{} `yaml:"TargetValue"`
	PredefinedMetricSpecification interface{} `yaml:"PredefinedMetricSpecification,omitempty"`
	CustomizedMetricSpecification interface{} `yaml:"CustomizedMetricSpecification,omitempty"`
}

// ScalingPolicyTargetTrackingConfiguration validation
func (resource ScalingPolicyTargetTrackingConfiguration) Validate() []error {
	errors := []error{}

	return errors
}
