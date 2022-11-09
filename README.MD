# SBIT-EMITTER

## Overview

This is a emitter service that send user request from sbit-api trough gRPC and emit those payload data to Kafka broker.   

- Please copy `config/.env.yaml` first to the root project and rename it to `.env.yaml` . Fill the field with your own environment parameters.    

- Running this service using command 
```bash
    go run main.go grpc
```