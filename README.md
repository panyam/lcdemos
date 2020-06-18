
# Installing Common Packages:

Docker
------

1. [Download and Install Docker](https://hub.docker.com/editions/community/docker-ce-desktop-mac/).

2. Run Docker.  This is needed to start the docker daemon which is then needed to create docker images in each of the projects.

# Cloud Specific Setup

## Google Cloud

0. [Demo App Setup Here](https://cloud.google.com/kubernetes-engine/docs/tutorials/hello-app).

1. [Install Google Cloud Tools](https://cloud.google.com/sdk/docs/quickstarts)

2. Install GC Kubectl

```
gcloud components install kubectl
```

3. Setup a GC Project from [Kubernetes Engine Page](https://console.cloud.google.com/projectselector/kubernetes).
