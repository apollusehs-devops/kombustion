package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// DirectoryConfigServiceAccountCredentials Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-directoryconfig-serviceaccountcredentials.html
type DirectoryConfigServiceAccountCredentials struct {
	AccountName     interface{} `yaml:"AccountName"`
	AccountPassword interface{} `yaml:"AccountPassword"`
}

// DirectoryConfigServiceAccountCredentials validation
func (resource DirectoryConfigServiceAccountCredentials) Validate() []error {
	errors := []error{}

	return errors
}
