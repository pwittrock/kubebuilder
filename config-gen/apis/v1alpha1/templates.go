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

import "text/template"

var (
	configTemplates = []*template.Template{
		// Template for the controller-manager namespace
		template.Must(template.New("namespace").Parse(`{{ if not .DisableCreateNamespace }}
apiVersion: v1
kind: Namespace
metadata:
  labels:
    control-plane: controller-manager
  name: {{ .Name }}-system
---
{{- end }}`)),

		// Template for the controller-manager
		template.Must(template.New("controller-manager").Parse(`{{ if not .DisableCreateManager }}
apiVersion: apps/v1
kind: Deployment
metadata:
  name: controller-manager
  namespace: {{ .Name }}-system
  labels:
    control-plane: controller-manager
spec:
  selector:
    matchLabels:
      control-plane: controller-manager
  replicas: 1
  template:
    metadata:
      labels:
        control-plane: controller-manager
    spec:
      containers:
      - command:
        - /manager
        args:
        - --enable-leader-election
        image: {{ .Image }}
        name: manager
        resources:
          limits:
            cpu: 100m
            memory: 30Mi
          requests:
            cpu: 100m
            memory: 20Mi
      terminationGracePeriodSeconds: 10
---
{{- if .EnableWebhooks }}
apiVersion: v1
kind: Service
metadata: 
  namespace: {{.Name}}-system
  name: webhook-service
  labels:
    control-plane: webhook
spec:
  ports:
  - port: 443
    targetPort: webhook-server
  selector:
    control-plane: controller-manager
---
{{- end}}
{{- if not .DisableAuthProxy }}
apiVersion: v1
kind: Service
metadata:
  namespace: {{ .Name }}-system
  name: metrics-service
  labels:
    control-plane: controller-manager
spec:
  ports:
  - name: https
    port: 8443
    targetPort: https
  selector:
    control-plane: controller-manager
---
{{ end }}{{ end }}`)),

		// Template for controller-manager RBAC role binding.
		// Note: the ClusterRole is generated by controller-gen from the code
		template.Must(template.New("rbac-manager").Parse(`{{- if not .DisableCreateRBAC }}
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: {{ .Name }}-manager-rolebinding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: {{ .Name }}-manager-role
subjects:
- kind: ServiceAccount
  name: default
  namespace: {{ .Name }}-system
---
{{ end }}`)),

		// Template for controller-manager RBAC leader election
		template.Must(template.New("rbac-leader-election").Parse(`{{- if not .DisableCreateRBAC }}
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: {{ .Name }}-leader-election-role
  namespace: {{ .Name }}-system
rules:
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - ""
  resources:
  - configmaps/status
  verbs:
  - get
  - update
  - patch
- apiGroups:
  - ""
  resources:
  - events
  verbs:
  - create
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: {{ .Name }}-leader-election-rolebinding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: {{ .Name }}-leader-election-role
subjects:
- kind: ServiceAccount
  name: default
  namespace: {{ .Name }}-system
---
{{ end }}`)),

		// Template for controller-manager RBAC metrics auth proxy
		template.Must(template.New("rbac-metrics-auth-proxy").Parse(`{{- if not .DisableCreateRBAC }}{{ if not .DisableAuthProxy}}
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: {{ .Name }}-proxy-role
rules:
- apiGroups: ["authentication.k8s.io"]
  resources:
  - tokenreviews
  verbs: ["create"]
- apiGroups: ["authorization.k8s.io"]
  resources:
  - subjectaccessreviews
  verbs: ["create"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: {{ .Name }}-proxy-rolebinding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: {{ .Name }}-proxy-role
subjects:
- kind: ServiceAccount
  name: default
  namespace: {{ .Name }}-system
---
{{ end }}{{ end }}`)),

		// Template for controller-manager prometheus service monitor
		template.Must(template.New("prometheus").Parse(`{{- if .EnablePrometheus }}
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  namespace: {{.Name}}-system
  name: controller-manager-metrics-monitor
  labels:
    control-plane: controller-manager
spec:
  endpoints:
    - path: /metrics
      port: https
  selector:
    matchLabels:
      control-plane: controller-manager
---
{{ end }}`)),

		// Template for the certification manager injection
		template.Must(template.New("cert-manager").Parse(`{{- if .EnableCertManager}}
# The following manifests contain a self-signed issuer CR and a certificate CR.
# More document can be found at https://docs.cert-manager.io
# WARNING: Targets CertManager 0.11 check https://docs.cert-manager.io/en/latest/tasks/upgrading/index.html for 
# breaking changes
apiVersion: cert-manager.io/v1alpha2
kind: Issuer
metadata:
  name: {{.Name}}-selfsigned-issuer
  namespace: {{.Name}}-system
spec:
  selfSigned: {}
---
apiVersion: cert-manager.io/v1alpha2
kind: Certificate
metadata:
  name: {{.Name}}-serving-cert
  namespace: {{.Name}}-system
spec:
  dnsNames:
  - {{.Name}}-webhook-service.{{.Name}}-system.svc
  - {{.Name}}-webhook-service.{{.Name}}-system.svc.cluster.local
  issuerRef:
    kind: Issuer
    name: selfsigned-issuer
  secretName: webhook-server-cert # this secret will not be prefixed, since it's not managed by kustomize
---
{{ end }}`)),
	}
)
