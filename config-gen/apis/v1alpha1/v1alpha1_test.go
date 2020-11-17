package v1alpha1_test

import (
	"testing"

	"sigs.k8s.io/kubebuilder/v2/config-gen/apis/v1alpha1"
	"sigs.k8s.io/kustomize/kyaml/fn/framework/frameworktestutil"
)

func TestNewCommand(t *testing.T) {
	test := frameworktestutil.ResultsChecker{Command: v1alpha1.NewCommand}
	test.Assert(t)
}
