#!/bin/bash

set -e
set -u
set -o pipefail

DOCKER_TAG=v0.1

MINIO_TAG=ascend-io/minio-gateway:$DOCKER_TAG

GCP_PROJECT=ascend-io-ops-artifacts

MINIO_GCR_TAG=gcr.io/$GCP_PROJECT/$MINIO_TAG

function build() {
  docker build -t $MINIO_TAG \
    -f Dockerfile.ascend \
    .
  docker tag $MINIO_TAG $MINIO_GCR_TAG
}

function push() {
  # Auth
  gcloud auth login
  gcloud config set project $GCP_PROJECT
  gcloud auth configure-docker

  # Minio push
  docker push $MINIO_GCR_TAG
}

build
push
