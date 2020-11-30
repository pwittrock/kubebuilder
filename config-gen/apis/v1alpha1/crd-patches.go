package v1alpha1

import (
	"github.com/markbates/pkger"
	"sigs.k8s.io/kustomize/kyaml/fn/framework"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

// CRDPatchTemplate returns the PatchTemplate for crd
func CRDPatchTemplate(kp *KubebuilderProject) framework.PT {
	return framework.PT{
		Dir: pkger.Dir("/config-gen/templates/patches/crd"),
		Selector: func() *framework.Selector {
			return &framework.Selector{
				Kinds: []string{"CustomResourceDefinition"},
				Filter: func(r *yaml.RNode) bool {
					m, _ := r.GetMeta()
					return kp.Spec.ConversionWebhooks[m.Name]
				},
			}
		},
	}
}
