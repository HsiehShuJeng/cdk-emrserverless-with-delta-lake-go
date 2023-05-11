//go:build !no_runtime_type_checking

package cdkemrserverlesswithdeltalake

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateEmrStudioTaggingExpert_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewEmrStudioTaggingExpertParameters(scope constructs.Construct, name *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

