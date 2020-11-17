package v1alpha1

import (
	"text/template"

	"sigs.k8s.io/kustomize/kyaml/fn/framework"
)

var (
	controllerManagerSelector = &framework.Selector{
		Kinds:            []string{"Deployment"},
		Namespaces:       []string{"{{ .Name }}-system"},
		Names:            []string{"controller-manager"},
		Labels:           map[string]string{"control-plane": "controller-manager"},
		TemplatizeValues: true,
	}

	patchTemplates = []framework.PatchTemplate{
		// Patch to enable the auth proxy in the controller-manager
		{
			Selector: controllerManagerSelector,
			Template: template.Must(template.New("controller-manager-auth-proxy-patch").Parse(`{{ if not .DisableAuthProxy}}
spec:
  template:
    spec:
      containers:
      - name: kube-rbac-proxy
        image: gcr.io/kubebuilder/kube-rbac-proxy:v0.5.0
        args:
        - "--secure-listen-address=0.0.0.0:8443"
        - "--upstream=http://127.0.0.1:8080/"
        - "--logtostderr=true"
        - "--v=10"
        ports:
        - containerPort: 8443
          name: https
      - name: manager
        args:
        - "--metrics-addr=127.0.0.1:8080"
        - "--enable-leader-election"
{{ end }}`)),
		},

		// Patch to enable the webhook server in the controller-manager
		{
			Selector: controllerManagerSelector,
			Template: template.Must(template.New("controller-manager-webhooks").Parse(`{{ if .EnableWebhooks }}
spec:
  template:
    spec:
      containers:
      - name: manager
        ports:
        - containerPort: 9443
          name: webhook-server
          protocol: TCP
        volumeMounts:
        - mountPath: /tmp/k8s-webhook-server/serving-certs
          name: cert
          readOnly: true
      volumes:
      - name: cert
        secret:
          defaultMode: 420
          secretName: webhook-server-cert
{{ end }}`)),
		},
	}
)
