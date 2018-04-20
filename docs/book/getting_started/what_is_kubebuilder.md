{% panel style="danger", title="Under Development" %}
This book is currently under development and does not reflect the current state of
the kubebuilder project!

Some of the APIs and libraries described in this book are proposals, and have not yet
been implemented!
{% endpanel %}

# What is Kubebuilder

Kubebuilder is an SDK for rapidly building and publishing Kubernetes APIs in the go language using the
canonical techniques that power Kubernetes.

In the spirit of modern web development frameworks such as *Ruby on Rails* and *SpringBoot*,
Kubebuilder provides a set of tools and libraries intended to simplify API development, and to
delight and empower developers.

Kubebuilder accomplishes this through providing:

* Tools to initialize *go* projects with the canonical set of libraries and their transitive dependencies
  necessary to build Kubernetes APIs.
* Tools to bootstrap new API definitions through writing scaffolding code, tests, and documentation.
* Simple, clean, high level libraries for invoking the Kubernetes APIs from go.
* Seamless integration of standard production logging and monitoring into API implementations.
* Tools to build and publish APIs as cluster addons or installable yaml declarations.
* Tools to build and publish API reference documentation with examples.
* Step by step guidance on how to use kubebuilder to develop your APIs.
