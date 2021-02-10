module sigs.k8s.io/kubebuilder/v3

go 1.15

require (
	github.com/cloudflare/cfssl v1.5.0
	github.com/go-logr/logr v0.3.0 // indirect
	github.com/gobuffalo/flect v0.2.2
	github.com/markbates/pkger v0.17.1
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.2
	github.com/spf13/afero v1.2.2
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	golang.org/x/tools v0.0.0-20200616195046-dc31b401abb5
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
	k8s.io/api v0.20.2 // indirect
	k8s.io/apiextensions-apiserver v0.20.1 // indirect
	k8s.io/apimachinery v0.20.2
	k8s.io/utils v0.0.0-20210111153108-fddb29f9d009 // indirect
	sigs.k8s.io/controller-runtime v0.8.2
	sigs.k8s.io/controller-tools v0.4.1
	sigs.k8s.io/kustomize/kyaml v0.10.2
	sigs.k8s.io/yaml v1.2.0
)

replace sigs.k8s.io/kustomize/kyaml => github.com/KnVerey/kustomize/kyaml v0.10.7-0.20210217003643-ddd01072985d
