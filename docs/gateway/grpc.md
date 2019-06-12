# MinIO GRPC Gateway [![Slack](https://slack.min.io/slack?type=svg)](https://slack.min.io)

MinIO GRPC Gateway allows you to serve arbitrary content via GRPC giving callers an Amazon S3-compatible API.

The theory is sometimes there is quite a bit of custom application logic in how paths map to the actual bits that are returned. It is very nice for people to be able to use native S3 clients built into most ecosystems / languages to interact with the datasources, but the mapping isn't 1-1. In addition, the permissions systems around access to specific resources may not be trivially accessible via static prolicies + sets of users. This GRPC gateway allows MinIO to have a shared-nothing architecture for processing the S3 requests, then forward to a user-provided GRPC service which can do whatever request processing is required to handle the application-specific URL transofrms and the like to return a byte strem.


TODO: Write the docs.