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

package main

import (
	"fmt"
	"os"

	"sigs.k8s.io/kubebuilder/v3/config-gen/apis/v1alpha1"
)

func main() {
	// Default to using the project file if no file is specified and
	// not being run as a function
	if os.Getenv("KUSTOMIZE_FUNCTION") != "true" && len(os.Args) <= 1 {
		if _, err := os.Stat("PROJECT"); err == nil {
			os.Args = append(os.Args, "PROJECT")
		}
	}

	cmd := v1alpha1.NewCommand()
	cmd.Use = "config-gen [CONFIG_FILE] [PATCHES...]"
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "\n%v\n", err)
		os.Exit(1)
	}
}
