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
	"bytes"
	"io"
	"path/filepath"

	"sigs.k8s.io/controller-tools/pkg/crd"
	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/rbac"
	"sigs.k8s.io/controller-tools/pkg/webhook"
	"sigs.k8s.io/kustomize/kyaml/errors"
	"sigs.k8s.io/kustomize/kyaml/kio"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

var _ kio.Filter = &ControllerGenFilter{}

// ControllerGenFilter generates resources using controller-gen
type ControllerGenFilter struct {
	*APIConfiguration
}

// Filter implements kio.Filter
func (cgr ControllerGenFilter) Filter(input []*yaml.RNode) ([]*yaml.RNode, error) {
	crdGen := genall.Generator(crd.Generator{})
	gens := genall.Generators{&crdGen}
	if !cgr.DisableCreateRBAC {
		rbacGen := genall.Generator(rbac.Generator{})
		gens = append(gens, &rbacGen)
	}
	if cgr.EnableWebhooks {
		webhookGen := genall.Generator(webhook.Generator{})
		gens = append(gens, &webhookGen)
	}

	b := BufferedGenerator{}
	rt, _ := gens.ForRoots(filepath.Join(cgr.Directory, "./..."))
	rt.OutputRules = genall.OutputRules{Default: &b}
	_ = rt.Run()

	// Parse the output
	n, err := (&kio.ByteReader{Reader: &b.Buffer}).Read()
	if err != nil {
		return nil, errors.WrapPrefixf(err, "failed to parse controller-gen output")
	}
	return append(n, input...), nil
}

type BufferedGenerator struct {
	bytes.Buffer
}

func (o *BufferedGenerator) Open(_ *loader.Package, _ string) (io.WriteCloser, error) {
	return o, nil
}

func (n BufferedGenerator) Close() error {
	return nil
}
