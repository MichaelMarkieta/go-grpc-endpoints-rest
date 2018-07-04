#### Tree

```
.
├── api_config.yaml	Cloud Endpoints configuration
├── client
│   └── main.go		Client application (gRPC)
├── cloudbuild.yaml	Cloud Container Builder configuration
├── customer
│   └── customer.proto	Protobuf RPC service definition
├── deployment.yaml	Google Kubernetes Engine deployment configuration
├── Dockerfile		Golang gRPC server container definition
├── Dockerfile_client	Golang gRPC client container definition
├── protoc
│   └── Dockerfile	Protobuf container definition
├── README.md
└── server
    └── main.go		API server (gRPC)
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
