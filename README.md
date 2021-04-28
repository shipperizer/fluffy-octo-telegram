# fluffy-octo-telegram

Simple grpc app implementing envoy CheckRequest Protocol for Client Authorization

## Build and deploy

Build setup is for multiarch support, a requirements for this is  [buildx](https://docs.docker.com/buildx/working-with-buildx/)

For `skaffold` integration i followed the suggestion [here](https://github.com/GoogleContainerTools/skaffold/tree/master/examples/custom-buildx) as ther eis no direct integration between `skaffold` and `buildx`


Images are pushed to `ghcr.io/shipperizer/fluffy-octo-telegram-grpc-app`, `k3s` cluster has a secret allowing it to pull them, see the snippet below in `deployments.yaml`

```
containers:
- image: ghcr.io/shipperizer/fluffy-octo-telegram-grpc-app
  name: fluffy-octo-telegram-grpc-app
  envFrom:
    - configMapRef:
        name: fluffy-octo-telegram-grpc-app
  name: fluffy-octo-telegram-grpc-app
  ports:
  - name: http
    containerPort: 8000
imagePullSecrets:
- name: regcred-github
```

### Kaniko

For `kaniko` builds, use the `--profile kaniko` modifier on `skaffold`, for this you will need an `Opaque` secret:

```
 echo '{"auths":{"ghcr.io":{"auth":"****************"}}}' | kubectl create secret generic regcred-github-kaniko --from-file=config.json=/dev/stdin
 ```

the profile is targeted at building on an `arm64` cluster only, if you need to use a different arch change `initImage` and `image` values

## ArgoCD

ArgoCD is used (together with ArgoCD image updater) to keep application up-to-date

see the `argocd.yaml` for extra informations
local setup will be described eventually in here, step by step

![ArgoCD](docs/argocd.png)
