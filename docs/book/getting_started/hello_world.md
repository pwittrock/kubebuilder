{% panel style="danger", title="Under Development" %}
This book is currently under development and does not reflect the current state of
the kubebuilder project!

Some of the APIs and libraries described in this book are proposals, and have not yet
been implemented!
{% endpanel %}

# Hello World

{% panel style="warning", title="Note on project structure" %}
Kubernete APIs require boilerplate code not shown here that is generated and maintained by kubebuilder.

Project structure will be initialized by running `kubebuilder init` and `kubebuilder create resource`.
More on this topic in *Project Creation and Structure*
{% endpanel %}

This chapter demonstrates writing a simple API definition that prints a log message after a resource
object is created, updated, or periodically reconciled.

A Kubernetes API has 3 components each of which is in a separate go package:

* The API schema definition - in a `go struct` declaration
* The API business logic implementation - in a `kb.Controller` object
* The main program - in a `go main function`

While the controller shown here only writes a log message, typical controllers
may read and write either Kubernetes objects or objects external to the cluster.

{% method %}
## Simple API Resource Definition {#hello-world-api}

The go struct declaration for the *Simple* API resource containing a single string field named *message*.

Note that additional boilerplate code is required but will be generated for the user by `kubebuilder generate`.

{% sample lang="go" %}
```go
// Note: This code lives under
// pkg/apis/*group*/*version*/simple_types.go

// Simple is a simple API that writes log messages
type Simple struct {
  // message is the message printed to the log
  Message string `json:"message"`
}
```
{% endmethod %}

{% method %}

## Simple Controller {#hello-world-controller}

The controller object that prints `Simple.message` to a log.  Reconcile
is called any time a *Simple* object is created / updated / deleted and also
periodically for each Simple object that exists (default interval 5 minutes).

{% sample lang="go" %}
```go
// Note: This code lives under
// pkg/controller/simple/controller.go

// create a new controller and register it with the sdk
func init() {
  // create a new controller that calls Reconcile
  c := &kb.ControllerMux{Reconcile: Reconcile}
	
  // respond to events
  kb.Handle(&v1beta1.Simple{}, c)
}

// Reconcile logs the message for a Simple resource
func Reconcile(k sdk.ReconcileKey) error {
  // create the instance to get
  s := &v1beta1.Simple{ObjectMeta: v1.ObjectMeta{
  	Name: k.Name, Namespace: k.Namespace,
  }}
  // get the instance
  if err := kb.Get(s); err != nil {
    if apierrors.IsNotFound(err) {
      // deleted, do nothing
      return nil
    }
    return err
  }
  log.Infof("%s", s.Message)
}
```
{% endmethod %}

{% method %}
## Simple Main {#hello-world-main}

The main program that imports known controller packages.

{% sample lang="go" %}
```go
// Note: This code lives under
// cmd/controller-manager/main.go

import(
  "flag"
  
  // register the Simple controller
  _ "pkg/controller/simple"
)

func main() {
  flag.Parse()
  log.Fatal(kb.ListenAndServe())
}
```
{% endmethod %}
