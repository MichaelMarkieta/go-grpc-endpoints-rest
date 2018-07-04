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

Work in this repo baed on steps from the [Exposing gRPC services as REST APIs using Cloud Endpoints](https://cloud.google.com/solutions/exposing-grpc-services-using-cloud-endpoints-pt1) tutorial.

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

#### Errors from ESP

```
INFO:Fetching an access token from the metadata service
INFO:Fetching the service config ID from the rollouts service
INFO:Fetching the service configuration from the service management service
nginx: [warn] Using trusted CA certificates file: /etc/nginx/trusted-ca-certificates.crt
10.28.33.1 - - [03/Jul/2018:01:17:44 +0000] "GET /v1/health HTTP/1.1" 401 337 "-" "kube-probe/1.8+"
10.28.33.1 - - [03/Jul/2018:01:17:49 +0000] "GET /v1/health HTTP/1.1" 401 337 "-" "kube-probe/1.8+"
10.28.33.1 - - [03/Jul/2018:01:17:54 +0000] "GET /v1/health HTTP/1.1" 401 337 "-" "kube-probe/1.8+"
10.28.33.1 - - [03/Jul/2018:01:17:59 +0000] "GET /v1/health HTTP/1.1" 401 337 "-" "kube-probe/1.8+"
[libprotobuf INFO external/servicecontrol_client_git/src/check_aggregator_impl.cc:244] Remove all entries of check aggregator.
[libprotobuf INFO external/servicecontrol_client_git/src/quota_aggregator_impl.cc:251] Remove all entries of check aggregator.
[libprotobuf INFO external/servicecontrol_client_git/src/report_aggregator_impl.cc:175] Remove all entries of report aggregator.
[libprotobuf INFO external/servicecontrol_client_git/src/report_aggregator_impl.cc:175] Remove all entries of report aggregator.
[libprotobuf INFO external/servicecontrol_client_git/src/quota_aggregator_impl.cc:251] Remove all entries of check aggregator.
[libprotobuf INFO external/servicecontrol_client_git/src/check_aggregator_impl.cc:244] Remove all entries of check aggregator.
```
