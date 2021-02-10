package v1alpha1

import (
	"github.com/markbates/pkger"
	"sigs.k8s.io/kustomize/kyaml/fn/framework"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

// CRDPatchTemplate returns the PatchTemplate for crd
func CRDPatchTemplate(kp *KubebuilderProject) framework.PatchTemplate {
	return &framework.ResourcePatchTemplate{
		TemplatesFn: framework.TemplatesFnFromDir(pkger.Dir("/config-gen/templates/patches/crd")),
		Selector: &framework.Selector{
			Kinds: []string{"CustomResourceDefinition"},
			MatchFn: func(r *yaml.RNode) bool {
				m, _ := r.GetMeta()
				return kp.Spec.ConversionWebhooks[m.Name]
			},
		},
	}
}
