# Changelog

### [1.0.3](https://www.github.com/shipperizer/fluffy-octo-telegram/compare/v1.0.2...v1.0.3) (2021-05-02)


### Bug Fixes

* exclude extServer.yaml from kustomize deployment, needs ot be done manually ([8a3ed8f](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/8a3ed8f72d6b02122f364711e0e9e8040c2535f9))

### [1.0.2](https://www.github.com/shipperizer/fluffy-octo-telegram/compare/v1.0.1...v1.0.2) (2021-04-29)


### Bug Fixes

* add log of grpc start ([a28cdc8](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/a28cdc8f3f13a8f429be2fad733a4f40dfbd2bfa))
* adjust deployment healthcheck ([da3105f](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/da3105f91ba391e7f8f7f863989f1274b19fae9b))
* disable tls healthcheck while i figure out how to do it ([3b4f53d](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/3b4f53dc08f0254ba676d933bf410178fdfffff9))
* use v1.0.0 image for a start ([aa77bf7](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/aa77bf78741072ac1d8551f84bd40c0ac8915726))

### [1.0.1](https://www.github.com/shipperizer/fluffy-octo-telegram/compare/v1.0.0...v1.0.1) (2021-04-29)


### Bug Fixes

* bad import for v2 ([6b9a1ae](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/6b9a1aeb0d6b3720e4038a671be35fba0d600faf))

## 1.0.0 (2021-04-28)


### ⚠ BREAKING CHANGES

* use kubectl to deploy argocd manifest
* move to camelCase for deployemnt files, adjust indentation
* adjust dockerfile label

### Features

* add cert manager as a dependency ([6a50b63](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/6a50b634d0fefe30328384e6cde144043e487305))
* add gh actions setup ([0217f46](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/0217f460e6645dc50c33ef6d1e61b07bfa6e72f2))
* deployment files and skaffold ([bff26e7](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/bff26e7f67d409d82a0ae8b588cc9756709b71c1))
* first version of auth server ([a5ddce1](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/a5ddce1a25d072f28a888332ff3070b36e7bed96))
* implement grpc healthcheck ([63b9922](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/63b9922a10e5cedce5f6adf5bd75409f1d44ef1f))
* initial setup for build ([700a749](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/700a7491a21ff3e6018d876ee34a6d87fc00716e))


### Bug Fixes

* adjust deps ([8aec228](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/8aec228fddce58a50cc999550239247687dce10c))
* adjust dockerfile label ([358bb09](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/358bb09c23dce62875237208ab818383e9606182))
* adjust repo on argocd file ([f57039c](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/f57039ca2e06f1f44632e01ff9534d529f4bc44f))
* bad rpc setup and wrong codes ([bf1e73c](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/bf1e73c186e5acb9d718b88fb2b69d37523b4511))
* move to camelCase for deployemnt files, adjust indentation ([76b83c9](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/76b83c99975a0a66bb18e4fba7e1587763180f0e))
* pointing at wrong secret ([c6ab03a](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/c6ab03acc3e472863a7001dfd7153a67469cef9d))
* put interface in a separate file ([ff467a1](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/ff467a1aa26bb0d0f92c84e642769249bf35cfb8))
* return httpResponse (wip) when authenticated ([fe6f149](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/fe6f149f5b5333d62e8cea95c5b61aeb8e2737ae))
* skaffold kaniko build needs extra args ([933c9ef](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/933c9ef4ff542ad890799eb67ecdac0f155b14a2))
* update argocd secrets ([3b2b5a0](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/3b2b5a072585cba32d8f736c1acacc93933a9396))
* update deps ([3c09c31](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/3c09c31a2a45c7225d919880d5535f9dd8933de1))
* use kubectl to deploy argocd manifest ([1c3a3af](https://www.github.com/shipperizer/fluffy-octo-telegram/commit/1c3a3aff37747cd19f6fe2a21cd08f50c0e68a32))
