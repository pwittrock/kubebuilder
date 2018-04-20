{% panel style="danger", title="Under Development" %}
This book is currently under development and does not reflect the current state of
the kubebuilder project!

Some of the APIs and libraries described in this book are proposals, and have not yet
been implemented!
{% endpanel %}

# Introduction

## Who is this for

#### Users of Kubernetes

This book well help users of Kubernetes develop a deep understanding of the fundamental concepts
behind how APIs are designed and how to best make use of them.  This knowledge will allow users
to better configure, manage and debug their workloads running in a Kubernetes cluster.

This book covers topics such as API versions semantics, how Spec / Status / ObjectMeta are
used in APIs, how the cluster performs self-healing, how objects are garbage collected and more..

#### Kubernetes API extension developers

This book provides a step-by-step guide for rapidly developing Kubernetes APIs in go.
Much of the content in this book is applicable regardless which language and platform
Kubernetes APIs are developed using.


#### Core Kubernetes developers

This book provides an overview of how the various pieces of the Kubernetes platform work
together to implement the Kubernetes container orchestrator.  Both new and veteran Kubernetes
contributors will benefit from the holistic picture presented in this book.

## Sharing this book

If you like this book, please use the buttons in the right of the top nav to share on Twitter,
Facebook and Google.

## Navigating this book

This section describes how to use the navigation elements of this book

##### Code navigation

Code samples may be either displayed to the side of the corresponding documentation, or inlined
immediately afterward.  This setting may be toggled using the split-screen icon in the left of the top nav.

##### Table of contents

The table of contents may be hidden using the hamburger icon in the left of the top nav.

##### OS / Language navigation

Some chapters have code snippets for multiple OS or Languages.  The chapters will display to OS
or Language selections in the right of the top nav.

## Getting a print version of this book

This book will soon be made available in print on Amazon for you to purchase.

## About the author

Phillip Wittrock is a Staff Software Engineer at Google working as a GKE and Kubernetes maintainer.
Phillip is a member of the Kubernetes Steering Committee, and has lead the following
Kubernetes Special Interest Groups: SIG cli, SIG release and SIG docs.

Phillip’s hobbies include debating how kubectl is pronounced, talking about Kubernetes API
at non-technical social events, and trying to make tools and libaries simple enough that even
he can use them.