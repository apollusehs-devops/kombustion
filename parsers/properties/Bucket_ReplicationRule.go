package properties

	import "fmt"

type Bucket_ReplicationRule struct {
	
	
	
	
	
	Id interface{} `yaml:"Id,omitempty"`
	Prefix interface{} `yaml:"Prefix"`
	Status interface{} `yaml:"Status"`
	SourceSelectionCriteria *Bucket_SourceSelectionCriteria `yaml:"SourceSelectionCriteria,omitempty"`
	Destination *Bucket_ReplicationDestination `yaml:"Destination"`
}

func (resource Bucket_ReplicationRule) Validate() []error {
	errs := []error{}
	
	
	
	
	
	if resource.Prefix == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Prefix'"))
	}
	if resource.Status == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Status'"))
	}
	if resource.Destination == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Destination'"))
	} else {
		errs = append(errs, resource.Destination.Validate()...)
	}
	return errs
}
