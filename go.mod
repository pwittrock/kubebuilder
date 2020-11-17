module sigs.k8s.io/kubebuilder/v2

go 1.15

require (
	github.com/gobuffalo/flect v0.2.2
	github.com/onsi/ginkgo v1.12.1
	github.com/onsi/gomega v1.10.1
	github.com/spf13/afero v1.2.2
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	golang.org/x/tools v0.0.0-20200616195046-dc31b401abb5
	k8s.io/apimachinery v0.19.4
	sigs.k8s.io/controller-runtime v0.6.3
	sigs.k8s.io/controller-tools v0.4.0
	sigs.k8s.io/kustomize/kyaml v0.9.5-0.20201114213312-b2ba82a0bdc2
	sigs.k8s.io/yaml v1.2.0
)

replace sigs.k8s.io/kustomize/kyaml => ../kustomize/kyaml
