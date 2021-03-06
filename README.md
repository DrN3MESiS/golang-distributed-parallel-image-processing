# Distributed Parallel Image Processing System

## Technologies
- CUDA
- gRPC
- REST API

## Language

- Golang
- C++

## Contributors

- Alan Enrique Maldonado Navarro
- Guillermo Gonzalez Mena

## Commands

- ./worker --controller tcp://localhost:40899 --worker-name Ciry --tags gpu,nvidia, assets, static

- ./worker --controller tcp://localhost:40899 --worker-name Miranda --tags gpu,nvidia, assets, static

- curl -u admin:password http://localhost:8080/login

- curl -F "data=boom.png" -H "Authorization: Bearer <token>" http://localhost:8080/upload

- curl -H "Authorization: Bearer <token>" http://localhost:8080/status

- curl -H "Authorization: Bearer <token>" http://localhost:8080/status/<Workername>

- curl -H "Authorization: Bearer <token>" http://localhost:8080/workloads/test

- curl -H "Authorization: Bearer <token>" http://localhost:8080/logout
