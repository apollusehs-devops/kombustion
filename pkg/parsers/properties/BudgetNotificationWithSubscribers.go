package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// BudgetNotificationWithSubscribers Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notificationwithsubscribers.html
type BudgetNotificationWithSubscribers struct {
	Notification interface{} `yaml:"Notification"`
	Subscribers  interface{} `yaml:"Subscribers"`
}

// BudgetNotificationWithSubscribers validation
func (resource BudgetNotificationWithSubscribers) Validate() []error {
	errors := []error{}

	return errors
}
