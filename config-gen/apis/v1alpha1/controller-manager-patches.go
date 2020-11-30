package v1alpha1

import (
	"github.com/markbates/pkger"
	"sigs.k8s.io/kustomize/kyaml/fn/framework"
)

// ControllerManagerPatchTemplate returns the PatchTemplate for controller-manager
func ControllerManagerPatchTemplate(kp *KubebuilderProject) framework.PT {
	return framework.PT{
		Dir: pkger.Dir("/config-gen/templates/patches/controller-manager"),
		Selector: func() *framework.Selector {
			return &framework.Selector{
				Kinds:            []string{"Deployment"},
				Namespaces:       []string{kp.Spec.Namespace},
				Names:            []string{"controller-manager"},
				Labels:           map[string]string{"control-plane": "controller-manager"},
				TemplatizeValues: true,
			}
		},
	}
}
