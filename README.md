Unary Mode | gRPC in Go Lang and Go Kit

What is Go kit? and why?
Go kit is the brain child of Peter Bourgon which he talked about during a FOSDEM and the Google Campus London meetup in 2015. As what he has envisioned, Go kit is a toolkit for building microservices in Go.

It provides you a set of standard libraries (packages) that are considered essentials when building microservices so you can focus on what’s important, your application’s business logic. Here is a list of packages Go kit provides

Authentication (Basic/Casbin/JWT)
Circuit Breaker (Hystrix/GoBreaker/HandyBreaker)
Logging: Provides an interface for structured logging.
Metrics: Provides an interface for service instrumentation. Has adapters for CloudWatch, Prometheus, Graphite, and more.
Rate Limiting
Service Discovery Utilities
Tracing (OpenCensus, OpenTracing, Zipkin)
Transport (HTTP/gRPC/ NATS/AWS Lambda, and more)
Other than being a toolkit, it also encourages good design principles for your application such as SOLID design principles, Domain-Driven Design (DDD), and Hexagonal architecture.Go kit's concept

Go kit's concept is based on three major components:

Services
Endpoints
Transports

Services (Business Logic)
This is where the core business logic of your API is located. In Go kit, each service method are converted into an Endpoint.

Endpoints
An endpoint represents a single RPC method. Each Service method converts into an Endpoint to make RPC style communication between servers and clients.

Transports
In a microservice architecture, a microservice often communicate through transports such as HTTP or gRPC. A single endpoint can be exposed by multiple transports.
