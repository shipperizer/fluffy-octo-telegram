#!/bin/sh
# This script demonstrates how to use `docker buildx` to build container
# images for the linux/amd64 and linux/arm64 platforms.  It creates a
# `docker buildx` builder instance when required.
#
# If you change the platforms, be sure to
#
#  (1) delete the buildx builder named `skaffold-builder`, and
#  (2) update the corresponding node-affinities in k8s/pod.yaml.

# The platforms to build.
platforms="linux/amd64,linux/arm,linux/arm64,Linux/riscv64"

app_name=$APP_NAME

# `buildx` uses named _builder_ instances configured for specific platforms.
# This script creates a `skaffold-builder` as required.
if ! docker buildx inspect skaffold-builder >/dev/null 2>&1; then
  docker buildx create --name skaffold-builder --platform $platforms
fi

# Building for multiple platforms requires pushing to a registry
# as the Docker Daemon cannot load multi-platform images.
if [ "$PUSH_IMAGE" = true ]; then
  args="--platform $platforms --push"
else
  args="--load"
fi

set -x      # show the command-line
docker buildx build --build-arg app_name=$app_name --builder skaffold-builder --tag $IMAGE $args "$BUILD_CONTEXT"
