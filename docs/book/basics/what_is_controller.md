{% panel style="danger", title="Under Development" %}
This book is currently under development and does not reflect the current state of
the kubebuilder project!

Some of the APIs and libraries described in this book are proposals, and have not yet
been implemented!
{% endpanel %}

# What is a Controller

Controllers implement APIs defined by *Resources*.  Controllers are
routines running in a Kubernetes cluster that watch both the resource API they implement as well
as related resource APIs to form a whole view of the cluster state.  Controllers reconcile each object's
(resource instance) desired state as declared in the Spec (e.g. 10 replicas of Pod running nginx)
with the state observed read from the APIs (e.g. 0 replicas of Pods running nginx).  Reconciliation is
done both in response to changes in cluster state, and periodically for each observed object.

**Kubernetes APIs and controllers have *level* based implementations to facilitate self-
healing and periodic reconciliation.**

## What is a Level Based API

The term *level-based* comes from interrupts hardware, where interrupts may be either *level-based* or *edge-based*.
This book does not go into the details of the hardware definitions of these terms.

Kubernetes defines a level-based API as implemented by reading the observed state of the system,
comparing it to the desired state declared in the object *Spec*, and *moving directly toward the
current desired state*.
 
This has a number of notable properties:

- reconciliation works directly towards the current desired state without having to completely pass through
  obsolete desired states
- when many events quickly occur that trigger a reconciliation for the same object, reconciliation will only be
  performed once or twice as only the observed and desired states are compared, not the events themselves.
- the system may trigger reconciliation periodically for objects without a specific event occurring.

Consider the following examples of level based API implementations.

**Example 1**: Skipping Obsolete States

A user creates a rollout for a new container image.  Shortly after starting the rollout, the user realizes
the containers are crash looping because they need to increase memory thresholds for the new image to
run.  The user updates the PodTemplate with the new memory limit and a new rollout is started.  In a
level based system, cluster will immediately start working towards the new target instead of trying
to complete the old rollout, whereas in an edge based system it would need to complete the first
(bad) rollout before starting the correct one.

**Example 2**: Batching Events

A user creates a Deployment with 1000 replicas.  The Deployment creates 1000 Pods and maintains a
Status field with the number of healthy Pods.  In a level based system, the controller doesn't
update the Status for each Pod (1000 writes), but instead batches updates together with
the number of observed healthy Pods during reconciliation.  In an edge baserd system, the
controller would respond to each individual Pod event with a Status update.

## Watching Events and Periodic Reconcile

The controller reconciliation between the declared desired state in the object and
the observed state of the cluster is triggered both by cluster events and periodically
for each object.

##### Watching Events

Controllers both watch for events on the resource they implement as well as related related resources.
For example the Deployment controller watches for *Deployment* events, *ReplicaSet* events and *Pod*
events.  This allows the Deployment to respond to change in cluster state, such as by continuing a
rolling update after newly started Pods become healthy by incrementally scaling down Pods with the old
template and scaling up Pods with the new template.

{% panel style="success", title="Mapping Events" %}
When the Deployment controller observes Pod events, it first maps them to the Deployment owning the Pod
and then inserts the Deployment key into a RateLimited queue to be reconciled.  This allows multiple
events for different Pods owned by the same Deployment to be batched together into a single reconcile.
{% endpanel %}

##### Periodic Reoncile

Each object is periodically reconciled even if no events are observed.  For this to be possible,
the reoncile function must only take the key of the object to reconcile, not a specific event.

## Generating Objects During Reconciliation

Many controllers generated new Kubernetes objects as part of a reconcile.  For example the
Deployment controller generates ReplicaSets, and the ReplicaSet controller generates Pods.
The controller ownership relationship between the generating and generated objects is
recorded both in an *OwnersReference* in the ObjectMeta of the generated objects and through
labels / selectors.  The labels / selectors allow the generating controller to find all of the
objects it has generated, and the *OwnersReference* confirms the relationship to address overlapping
or modified labels.

## Writing Status Back to Objects

Controllers are run asynchronously, meaning that the user will not get a status update
in response to `applying` (or creating, updating, patching) an object, since the controller
will not have reconciled the state.  In order to communicate status back to the user,
controllers write to the object *Status* field.

{% panel style="info", title="Status" %}
The controller will keep Status up-to-date both in response to user initiated events, but also
in response to non-user initiated events, such as Node failures.
{% endpanel %}
