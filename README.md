
# Installing Common Packages:

DB Adapters
-----------

Postgres:
=========

```
brew install libpq
```

Python
------

Install Python2
===============

Sad but Needed for most gcloud commandline tools - WTF?

```
pyenv install -f 2.7.17
```

Install Python 3.8+
===================

```
pyenv install -f 3.8.2
```

Docker
------

1. [Download and Install Docker](https://hub.docker.com/editions/community/docker-ce-desktop-mac/).

2. Run Docker.  This is needed to start the docker daemon which is then needed to create docker images in each of the projects.

# Environment Specific Setup

## Minikube

1. [Install Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)

2. 

## Google Cloud

0. [Demo App Setup Here](https://cloud.google.com/kubernetes-engine/docs/tutorials/hello-app).

1. [Install Google Cloud Tools](https://cloud.google.com/sdk/docs/quickstarts)

2. Install GC Kubectl

```
gcloud components install kubectl
```

3. Setup a GC Project from [Kubernetes Engine Page](https://console.cloud.google.com/projectselector/kubernetes).

4. Setup a Cluster:

Our clusters following a naming convention - `lcdemos-<system_name>-<version>`.

Eg the cluster for v1 of the Twitter usecase would be - `lcdemos-twitter-v1`.

To create the cluster:

```
PROJECT_ID=<YOUR_GOOGLE_CLOUD_PROJECT_ID>
CLUSTER_NAME=lcdemos_<casename>_<version>

# Set the zone for our project
gcloud config set compute/zone us-west1-a

# create the cluster - this will take a while
gcloud container clusters create ${CLUSTER_NAME}
```
