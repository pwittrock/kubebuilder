/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"io/ioutil"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/kustomize/kyaml/errors"
)

// KubebuilderProject implements the API for generating configuration
type KubebuilderProject struct {
	metav1.TypeMeta   `json:",inline" yaml:",omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Spec is the configuration spec defining what configuration should be produced.
	Spec KubebuilderProjectSpec `json:"spec,omitempty" yaml:"spec,omitempty"`

	// Status is the configuration status defined at runtime.
	Status KubebuilderProjectStatus `json:"status,omitempty" yaml:"status,omitempty"`

	// Project is for the field name used by the PROJECT file (config), instead of spec.
	// Use Spec unless in the PROJECT file.
	Project *KubebuilderProjectSpec `json:"config,omitempty" yaml:"config,omitempty"`
}

var (
	// NamespaceComponent will create the namespace for the controller-manager
	NamespaceComponent = "namespace"

	// ControllerManagerComponent will create the controller-manager Deployment
	ControllerManagerComponent = "controller-manager"

	// WebhooksComponent will create the webhook configurations
	WebhooksComponent = "webhooks"

	// CRDsComponent will create the CRDs
	CRDsComponent = "crds"

	// RBACComponent will create RBAC rules
	RBACComponent = "rbac"

	// CertManagerComponent will create the Issuer and Certificate resources
	// for CertManager to inject the certificates
	CertManagerComponent = "cert-manager"

	// PrometheusComponent will create the ServiceMonitor resource
	PrometheusComponent = "prometheus"
)

// getDefaultComponents returns the set of components that are created by default
func getDefaultComponents() map[string]bool {
	return map[string]bool{
		NamespaceComponent:         true,
		ControllerManagerComponent: true,
		RBACComponent:              true,
		CRDsComponent:              true,
	}
}

// Enabled returns true if the component is enabled
func (a *KubebuilderProjectSpec) Enabled(component string) bool {
	return a.Components[component]
}

// ComponentConfigEnabled returns true if component config is being used
func (a *KubebuilderProjectSpec) ComponentConfigEnabled() bool {
	return a.ComponentConfigFilepath != ""
}

// KubebuilderProjectSpec defines the desired configuration to be generated
type KubebuilderProjectSpec struct {
	// Directory is the kubebuilder directory containing the code
	Directory string `json:"projectDirectory" yaml:"projectDirectory"`

	// Name is the name of project and used to generate the component and role names
	// Defaults to metadata.name
	Name string `json:"projectName" yaml:"projectName"`

	// Namespace is the namespace to run the project in -- defaults to projectName-system
	Namespace string `json:"namespace" yaml:"namespace"`

	// Image is the container image to run in the controller-manager
	Image string `json:"image" yaml:"image"`

	// Components is a map of components to enable or disable
	Components map[string]bool `json:"components" yaml:"components"`

	// DisableAuthProxy if set to true will disable the auth proxy
	DisableAuthProxy bool `json:"disableAuthProxy,omitempty" yaml:"disableAuthProxy,omitempty"`

	// ConversionWebhooks is a map of kinds to enable conversion webhooks for
	ConversionWebhooks map[string]bool `json:"conversionWebhooks,omitempty" yaml:"conversionWebhooks,omitempty"`

	// Development contains development options
	Development DevelopmentOptions `json:"developmentOptions,omitempty" yaml:"developmentOptions,omitempty"`

	ComponentConfigFilepath string `json:"componentConfigFilepath,omitempty" yaml:"componentConfigFilepath,omitempty"`
}

// KubebuilderProjectStatus is runtime status for the api configuration
type KubebuilderProjectStatus struct {
	CertCA string

	CertKey string

	ComponentConfigString string
}

// DevelopmentOptions defines options for development installation
type DevelopmentOptions struct {
	// GenerateCert will cause a self signed certificate to be generated and injected
	// into the Webhook caBundles.
	GenerateCert bool `json:"generateCert,omitempty" yaml:"generateCert,omitempty"`

	// CertDuration sets the duration for the cert
	CertDuration time.Duration `json:"certDuration,omitempty" yaml:"certDuration,omitempty"`
}

// Default defaults the values
func (kp *KubebuilderProject) Default() error {
	// merge project.yaml into fc
	if kp.Project != nil {
		// For compabitility with the PROJECT file format
		kp.Spec = *kp.Project
	}

	// Validate the input
	if kp.Spec.Name == "" {
		if kp.Name == "" {
			return errors.Errorf("must specify Kubebuilder projectName field")
		}
		kp.Spec.Name = kp.Name
	}
	if kp.Spec.Namespace == "" {
		if kp.Namespace != "" {
			kp.Spec.Namespace = kp.Namespace
		}
		kp.Spec.Namespace = kp.Spec.Name + "-system"
	}
	if kp.Spec.Image == "" {
		return errors.Errorf("must specify Kubebuilder image field")
	}
	if kp.Spec.Development.CertDuration == 0 {
		d := time.Hour * 1
		kp.Spec.Development.CertDuration = d
	}

	if kp.Spec.Directory == "" {
		kp.Spec.Directory = "./..."
	}

	if kp.Spec.Components == nil {
		kp.Spec.Components = map[string]bool{}
	}
	for k, v := range getDefaultComponents() {
		if _, found := kp.Spec.Components[k]; !found {
			kp.Spec.Components[k] = v
		}
	}

	if kp.Spec.ComponentConfigFilepath != "" {
		b, err := ioutil.ReadFile(kp.Spec.ComponentConfigFilepath)
		if err != nil {
			return err
		}
		kp.Status.ComponentConfigString = string(b)
	}

	return nil
}
