# Config-gen

This is a prototype of using a kustomize function to generate configuration.

To test it:

```sh
go build -o ~/go/bin/config-gen ./config-gen
```

## Usage Options

### With PROJECT file

Add the config gen fields to your project file

```yaml
# PROEJCT
metadata:
  name: project-name # name used to generated resource names and namespaces
spec:
  image: pwittrock/simple # controller-manager image to run
...
```

```sh
# from a kubebuilder project
config-gen
```

### With config file

Create a config.yaml

```yaml
# config.yaml
apiVersion: kubebuilder.sigs.k8s.io
kind: APIConfiguration
metadata:
  name: project-name
spec:
  image: example/simple:latest
```

```sh
# from a kubebuilder project
config-gen config.yaml
```

### With patch overrides

```yaml
# config.yaml
apiVersion: kubebuilder.sigs.k8s.io
kind: APIConfiguration
metadata:
  name: project-name
spec:
  image: example/simple:latest
```

```yaml
# patch1.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: controller-manager
  namespace: project-name-system
spec:
...
```

```sh
# from a kubebuilder project
config-gen config.yaml patch1.yaml
```

### From kustomize

config-gen may be run as a Kustomize plugin using kustomize