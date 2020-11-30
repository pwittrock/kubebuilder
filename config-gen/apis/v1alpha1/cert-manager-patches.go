package v1alpha1

import (
	"github.com/markbates/pkger"
	"sigs.k8s.io/kustomize/kyaml/fn/framework"
)

// CertManagerPatchTemplate returns the PatchTemplate for cert-manager
func CertManagerPatchTemplate(kp *KubebuilderProject) framework.PT {
	return framework.PT{
		Dir: pkger.Dir("/config-gen/templates/patches/cert-manager"),
		Selector: func() *framework.Selector {
			return &framework.Selector{
				Kinds: []string{
					"CustomResourceDefinition",
					"ValidatingWebhookConfiguration",
					"MutatingWebhookConfiguration",
				},
			}
		},
	}
}
