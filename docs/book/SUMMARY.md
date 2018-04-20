# Building Kubernetes APIs with Kubebuilder

* [Introduction](README.md)

### Getting Started

* [Why Kubernetes APIs](getting_started/why_kubernetes.md)
* [What is Kubebuilder](getting_started/what_is_kubebuilder.md)
* [Installation and Setup](getting_started/installation_and_setup.md)
* [Hello World](getting_started/hello_world.md)

### Basics

* [Project Creation and Structure](basics/project_creation_and_structure.md)
* Resource Fundamentals
    * [What is a Resource](basics/what_is_resource.md)
    * [Simple Resource Example](basics/simple_resource.md)
* Controller Fundamentals
    * [What is a Contoller](basics/what_is_controller.md)
    * [Simple Controller Example](basics/simple_controller.md)
* Cmd Fundamentals
    * [Registering Controllers and Resources](basics/register_controllers_resources.md)
* [Development Workflow](basics/development_workflow.md)

### Beyond the Basics

* Resources
    * [Configuring Generated Code](beyond_the_basics/generated_code.md)
    * [Handling Read / Write Failures](beyond_the_basics/handling_failures.md)
    * [Non-Namespaced Objects](beyond_the_basics/generated_code.md)
    * [Multi Version Support](beyond_the_basics/resource_versioning.md)
    * [Atomic and Associative Fields](beyond_the_basics/atomic_and_associative_fields.md)
    * [Validation using OpenAPI](beyond_the_basics/schema_validation_with_openapi.md)
    * [Validation using Webhooks](beyond_the_basics/additional_validation_with_webhooks.md)
    * [Defaulting and Canonicalization using Webhooks](beyond_the_basics/defaulting_and_canonicalization_with_webhooks.md)
    * [Debugging Resources](beyond_the_basics/debugging_resources.md)

* Controllers
    * [Generating Unique Object Names](beyond_the_basics/generating_unique_object_names.md)
    * [Finalizers and Garbage Collection](beyond_the_basics/finalizers_and_garbage_collection.md)
    * [Events, Shared Informers and RateLimiting Queues](beyond_the_basics/events_shared_informers_and_ratelimiting_queues.md)
    * [Logging and Publishing Events](beyond_the_basics/logging_and_publishing_events.md)
    * [Debugging Controllers](beyond_the_basics/debugging_controllers.md)
    
* Controller Manager
    * [Initializing and Wiring Shared Dependencies](beyond_the_basics/initializing_and_wiring_dependencies.md)
    * [Installing CRDs](beyond_the_basics/installing_crds.md)

### Publishing and Running in Production

* [Integration Testing](publishing_and_running_in_production/integration_testing.md)
* [Monitoring and Alerting](publishing_and_running_in_production/monitoring_and_alerting.md)
* [Build and Release](publishing_and_running_in_production/build_and_release.md)
* [Packaging and Publishing](publishing_and_running_in_production/package_and_publishing.md)

### Raw Extension Points

* [CustomResourceDefinitions]()
* [Annotations as Virtual Fields]()
* [Webhooks]()
    * [Validating]()
    * [Mutating]()
* [CRUD APIs]()
    * [Get, List, Watch]()
    * [Create, Update, Patch]()
* [Init Containers]()


### Reference

* [Resource Definitions]()
* [Controller Definitions]()