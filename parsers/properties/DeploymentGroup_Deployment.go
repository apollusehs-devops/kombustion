package properties

	import "fmt"

type DeploymentGroup_Deployment struct {
	
	
	
	Description interface{} `yaml:"Description,omitempty"`
	IgnoreApplicationStopFailures interface{} `yaml:"IgnoreApplicationStopFailures,omitempty"`
	Revision *DeploymentGroup_RevisionLocation `yaml:"Revision"`
}

func (resource DeploymentGroup_Deployment) Validate() []error {
	errs := []error{}
	
	
	
	if resource.Revision == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Revision'"))
	} else {
		errs = append(errs, resource.Revision.Validate()...)
	}
	return errs
}
