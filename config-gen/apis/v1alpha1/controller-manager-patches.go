package v1alpha1

import (
	"github.com/markbates/pkger"
	"sigs.k8s.io/kustomize/kyaml/fn/framework"
)

// ControllerManagerPatchTemplate returns the PatchTemplate for controller-manager
func ControllerManagerPatchTemplate(kp *KubebuilderProject) framework.PatchTemplate {
	return &framework.ResourcePatchTemplate{
		TemplatesFn: framework.TemplatesFnFromDir(pkger.Dir("/config-gen/templates/patches/controller-manager")),
		Selector: &framework.Selector{
			Kinds:        []string{"Deployment"},
			Namespaces:   []string{"{{ .Spec.Namespace }}"},
			Names:        []string{"controller-manager"},
			Labels:       map[string]string{"control-plane": "controller-manager"},
			TemplateData: kp,
		},
	}
}
