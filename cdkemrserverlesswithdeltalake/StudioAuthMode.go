package cdkemrserverlesswithdeltalake


// What kind of authentication the Studio uses.
type StudioAuthMode string

const (
	// the Studio authenticates users using AWS SSO.
	StudioAuthMode_AWS_SSO StudioAuthMode = "AWS_SSO"
	// the Studio authenticates users using AWS IAM.
	StudioAuthMode_AWS_IAM StudioAuthMode = "AWS_IAM"
)

