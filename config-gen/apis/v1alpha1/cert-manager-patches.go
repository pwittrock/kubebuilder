package v1alpha1

import (
	"github.com/markbates/pkger"
	"sigs.k8s.io/kustomize/kyaml/fn/framework"
)

// CertManagerPatchTemplate returns the PatchTemplate for cert-manager
func CertManagerPatchTemplate(kp *KubebuilderProject) framework.PatchTemplate {
	return &framework.ResourcePatchTemplate{
		TemplatesFn: framework.TemplatesFnFromDir(pkger.Dir("/config-gen/templates/patches/cert-manager")),
		Selector: &framework.Selector{
			Kinds: []string{
				"CustomResourceDefinition",
				"ValidatingWebhookConfiguration",
				"MutatingWebhookConfiguration",
			},
		},
	}
}
