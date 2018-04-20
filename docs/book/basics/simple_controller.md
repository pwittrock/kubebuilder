{% panel style="danger", title="Under Development" %}
This book is currently under development and does not reflect the current state of
the kubebuilder project!

Some of the APIs and libraries described in this book are proposals, and have not yet
been implemented!
{% endpanel %}

# Simple Resource Example

This chapter walks through a simple Controller implementation.

This is a simple controller example for the ContainerSet API shown in *Simple Resource Example*.

The controller reads the 

> $ kubebuilder create resource --group workloads --version v1beta1 --kind ContainerSet
> pkg/controller/containerset/controller.go

{% method %}
## Setup

The controller is setup in the package `init` function.  Any errors during setup should be
be returned when starting the controller manager, not in the init function.

- Create a new `ControllerMux` with the Reconcile function specified.
- Watch for ContainerSet events and reconcile the corresponding ContainerSet object
- Watch for Deployment events and reconcile the Owner object if the reference has "controller: true",
  and the Owner type is a ContainerSet

{% sample lang="go" %}
```go
func init() {
  c := &kb.ControllerMux{Reconcile: Reconcile}

  kb.Handle(&v1beta1.ContainerSet{}, c)

  kb.Handle(kb.DispatchToOwner{
  	Generated: &v1.Deployment{},
  	Owner: &v1beta1.ContainerSet{},
  	Controller: true,
  }}, c)
}
```
{% endmethod %}

{% method %}
## Implementation

The controller is implemented in the `Reconcile` function.  This function takes the namespace/name
key of the ContainerSet object to reconcile.  It then reads the ContainerSet object, checks
if a matching Deployment already exists, and either creates or updates the Deployment.

Finally the controller updates the Status on the ContainerSet.  Because the Deployment and ContainerSet
cannot be written in a single transaction, in the event the Status update fails the controller will
need to set the Status during the following Reconcilation.

Note that if the Deployment is deleted or changed by some other actor in the system, the controller
will receive and event and recreate / update the Deployment.

{% sample lang="go" %}

```go
func Reconcile(k sdk.ReconcileKey) error {
  s := &v1beta1.ContainerSet{ObjectMeta: v1.ObjectMeta{
  	Name: k.Name, Namespace: k.Namespace,
  }}
  if err := kb.Get(s); err != nil {
    if apierrors.IsNotFound(err) {
      return nil
    }
    return err
  }
  
  // Create / Update Deployment
  d := &v1.Deployment{
  	ObjectMeta: v1.ObjectMeta{
      	Name: k.Name, Namespace: k.Namespace,
    },
    Spec: v1.DeploymentSpec{...},
  }
  kb.SetOwnerReference(d, s)
  
  err := kb.Get(d)
  if err != nil && !apierrors.IsNotFound(err) {
    return err
  }
  if apierrors.IsNotFound(err) {
      if err := kb.Create(d); err != nil {
        return err
      }  	
  } else {
      d.Spec = Spec: v1.DeploymentSpec{...}
      kb.SetOwnerReference(d, s)
      if kb.Update(d); err != nil {
        return err
      }  
  }
  
  // Update ContainerSet Status
  s.Status.HealthyReplicas = d.Status.ReadyReplicas
  if err := kb.Update(s); err != nil {
      return err
  }
  return nil
}
```
{% endmethod %}


