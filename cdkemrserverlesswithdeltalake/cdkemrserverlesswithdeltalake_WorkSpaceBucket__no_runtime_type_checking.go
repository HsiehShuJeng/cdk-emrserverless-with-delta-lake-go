//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// A construct for the quick demo of EMR Serverless.
package cdkemrserverlesswithdeltalake

// Building without runtime type checking enabled, so all the below just return nil

func validateWorkSpaceBucket_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewWorkSpaceBucketParameters(scope constructs.Construct, name *string, props *WorkSpaceBucketProps) error {
	return nil
}

