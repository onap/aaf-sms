#!/bin/bash
set -e
DIRNAME=`dirname $0`
DOCKER_BUILD_DIR=`cd $DIRNAME/; pwd`
echo "DOCKER_BUILD_DIR=${DOCKER_BUILD_DIR}"
cd ${DOCKER_BUILD_DIR}

BUILD_ARGS="--no-cache"
ORG="onap"
VERSION="4.0.0"
PROJECT="aaf"
IMAGE="smsquorumclient"
DOCKER_REPOSITORY="localhost:5000"
IMAGE_NAME="${DOCKER_REPOSITORY}/${ORG}/${PROJECT}/${IMAGE}"
TIMESTAMP=$(date +"%Y%m%dT%H%M%S")

if [ $HTTP_PROXY ]; then
    BUILD_ARGS+=" --build-arg HTTP_PROXY=${HTTP_PROXY}"
fi
if [ $HTTPS_PROXY ]; then
    BUILD_ARGS+=" --build-arg HTTPS_PROXY=${HTTPS_PROXY}"
fi
#Need to create makefile
function generate_binary {
    pushd ../src/quorumclient
    make build
    popd
    cp ../target/quorumclient .
}

function copy_certificates {
    cp ../src/sms/certs/aaf_root_ca.cer .
}

function cleanup {
    rm quorumclient
    rm aaf_root_ca.cer
}

function build_image {
    echo "Start build docker image: ${IMAGE_NAME}"
    docker build ${BUILD_ARGS} -t ${IMAGE_NAME}:latest -f quorumdockerfile .
}

function push_image_tag {
    TAG_NAME=$1
    echo "Start push ${TAG_NAME}"
    docker tag ${IMAGE_NAME}:latest ${TAG_NAME}
    docker push ${TAG_NAME}
}

function push_image {
    echo "Start push ${IMAGE_NAME}:latest"
    docker push ${IMAGE_NAME}:latest

    push_image_tag ${IMAGE_NAME}:${VERSION}-SNAPSHOT-latest
}

generate_binary
copy_certificates
build_image
push_image
cleanup
