package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// InstanceGroupConfigSimpleScalingPolicyConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-simplescalingpolicyconfiguration.html
type InstanceGroupConfigSimpleScalingPolicyConfiguration struct {
	AdjustmentType    interface{} `yaml:"AdjustmentType,omitempty"`
	CoolDown          interface{} `yaml:"CoolDown,omitempty"`
	ScalingAdjustment interface{} `yaml:"ScalingAdjustment"`
}

// InstanceGroupConfigSimpleScalingPolicyConfiguration validation
func (resource InstanceGroupConfigSimpleScalingPolicyConfiguration) Validate() []error {
	errors := []error{}

	return errors
}
