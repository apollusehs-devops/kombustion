package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// SizeConstraintSetSizeConstraint Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sizeconstraintset-sizeconstraint.html
type SizeConstraintSetSizeConstraint struct {
	ComparisonOperator interface{} `yaml:"ComparisonOperator"`
	Size               interface{} `yaml:"Size"`
	TextTransformation interface{} `yaml:"TextTransformation"`
	FieldToMatch       interface{} `yaml:"FieldToMatch"`
}

// SizeConstraintSetSizeConstraint validation
func (resource SizeConstraintSetSizeConstraint) Validate() []error {
	errors := []error{}

	return errors
}
