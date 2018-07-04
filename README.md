#### Tree

```
.
├── api_config.yaml
├── client
│   └── main.go
├── cloudbuild.yaml
├── customer
│   └── customer.proto
├── deployment.yaml
├── Dockerfile
├── Dockerfile_client
├── google
│   └── api
│       ├── annotations.proto
│       └── http.proto
├── protoc
│   └── Dockerfile
├── README.md
└── server
    └── main.go
```

#### Developer notes

```
// set environment variables
PROJECT_ID=PROJECT_ID
ZONE=ZONE
CLUSTER=CLUSTER

// set gcloud config
gcloud config set project $PROJECT_ID
gcloud config set compute/zone $ZONE

// build a base docker image from golang:alpine and install protoc from source, saving us time in subsequent steps
docker build -t protoc protoc
docker tag protoc gcr.io/$PROJECT_ID/protoc
docker push gcr.io/$PROJECT_ID/protoc

// deploy using cloudbuild
gcloud container builds submit --config=cloudbuild.yaml --substitutions _ZONE=$ZONE,_CLUSTER=$CLUSTER .
```
