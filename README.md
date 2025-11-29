# Combination Service

[![release](https://img.shields.io/github/v/release/qkitzero/combination-service?logo=github)](https://github.com/qkitzero/combination-service/releases)
[![test](https://github.com/qkitzero/combination-service/actions/workflows/test.yml/badge.svg)](https://github.com/qkitzero/combination-service/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/qkitzero/combination-service/graph/badge.svg)](https://codecov.io/gh/qkitzero/combination-service)
[![Buf CI](https://github.com/qkitzero/combination-service/actions/workflows/buf-ci.yaml/badge.svg)](https://github.com/qkitzero/combination-service/actions/workflows/buf-ci.yaml)

- Microservices Architecture
- gRPC
- gRPC Gateway
- Buf ([buf.build/qkitzero-org/combination-service](https://buf.build/qkitzero-org/combination-service))
- Clean Architecture
- Docker
- Test
- Codecov
- Cloud Build
- Cloud Run

```mermaid
flowchart TD
    subgraph gcp[GCP]
        secret_manager[Secret Manager]

        subgraph cloud_build[Cloud Build]
            build_combination_service(Build combination-service)
            push_combination_service(Push combination-service)
            deploy_combination_service(Deploy combination-service)

            build_combination_service_gateway(Build combination-service-gateway)
            push_combination_service_gateway(Push combination-service-gateway)
            deploy_combination_service_gateway(Deploy combination-service-gateway)
        end


        subgraph artifact_registry[Artifact Registry]
            combination_service_image[(combination-service image)]
            combination_service_gateway_image[(combination-service-gateway image)]
        end

        subgraph cloud_run[Cloud Run]
            combination_service(Combination Service)
            combination_service_gateway(Combination Service Gateway)
        end
    end

    subgraph external[External]
        auth_service(Auth Service)
        combination_db[(Combination DB)]
    end

    build_combination_service --> push_combination_service --> combination_service_image
    build_combination_service_gateway --> push_combination_service_gateway --> combination_service_gateway_image

    combination_service_image --> deploy_combination_service --> combination_service
    combination_service_gateway_image --> deploy_combination_service_gateway --> combination_service_gateway

    secret_manager --> deploy_combination_service
    secret_manager --> deploy_combination_service_gateway

    combination_service_gateway --> combination_service
    combination_service --> combination_db
    combination_service --> auth_service
```
