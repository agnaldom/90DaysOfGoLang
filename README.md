# 90DaysOfGoLang

This respository is used to document my journey on getting a better understinf of Golang Developer, I Will be starting this journey on the 13rd of april 2022, but the idea is that it will take uns 90 days, which happens until then.

The reason for documenting these days is os that others can take something out of it and also improve the features.

## Begin Concepts

This course has syntax training and exercises. Designer for anyone who has programming experience but has little to no experience with Go.

1.0 - Hello World
2.0 - Variables
3.0 - Strings
4.0 - Struct
5.0 - Type Conversion
6.1 - Pointers PT.1
6.2 - Pointers PT.2
6.3 - Pointers PT.3
6.4 - Pointers PT.4
7.0 - Literal Struct
8.1 - Constants PT.1
8.2 - Constants PT.2
8.3 - Constants PT.3
8.4 - Constants PT.4
9.0 - Blank Identifier
10.0 - Function Idioms
11.0 - Literal Functions
12.0 - Short Variables Declaration Operator
13.0 - Array Basics
14.0 - Array Iterarion
15.0 - Array Mismatch

## Advanced Concetps

This course teaches how to write better, more idiomatic and performant code in Go with a focus on micro-level engineering decisions. The course begins with the a focus on Go internals that are critical to understanding the core tradeoffs the language makes on readability, simplicity and performance.

### 01. Design Guidelines

1.1 - Prepare Your Mind
1.2 - Productivity vs. Performance
1.3 - Correctness vs. Performance
1.4 - Code Reviews
1.5 - If Performance Matters

### 02. Memory & Data Semantics

2.1 - Variables
2.2 - Struct Types
2.3.1 - Pointers-Part 1 (Pass by Values)
2.3.2 - Pointers-Part 2 (Sharing Data)
2.3.3 - Pointers-Part 3 (Escape Analysis)
2.3.4 - Pointers-Part 4 (Stack Growth)
2.3.5 - Pointers-Part 5 (Garbage Colletion)
2.4 - Constants

### 03. Data Structures

3.1 - Arrays-Part 1 (Mechanical Sympathy)
3.2.1 - Arrays-Part 2 (Semantics)
3.2.2 - Arrays-Part 3 (Range Mechanics)
3.3.1 - Arrays-Part 1 (Declare, Length & Reference Types)
3.3.2 - Arrays-Part 2 (Appending Slices)
3.3.3 - Arrays-Part 3 (Taking Slices of Slices)
3.3.4 - Slices-Part 4 (Slices & References)
3.3.5 - Slices-Part 5 (Strings & Slices)
3.3.6 - Slices-Part 6 (Range Mechanics)
3.4 - Maps

### 04. Decoupling

4.1.1 - Methods-Part 1 (Value & Pointer Semantics)
4.1.2 - Methods-Part 2 (Function/Method Variables)
4.2.1 - Interfaces-Part 1 (Polymorphism)
4.2.2 - Interfaces-Part 1 (Method Sets & Address of Value)
4.2.3 - Interfaces-Part 3 (Storage by Value)
4.2.4 - Interfaces-Part 4 (Type Assertion)
4.3 - Embedding
4.4 - Exporting

### 05. Composition

5.1 - Grouping Types
5.2.1 - Decoupling-Part 1
5.3.1 - Conversion & Assertions-Part 1
5.3.2 - Conversion & Assertions-Part 2
5.4 - Interface Pollution
5.5 - Mocking

### 06. Error Handling

6.1 - Default Error Values
6.2 - Error Variables
6.3 - Types as Context
6.4 - Behavior as Context
6.5 - Find the Bug
6.6 - Wrapping Errors

### 07. Packaging

7.1 - Language Mechanics & Design Guidelines
7.2 - Package-Oriented Design

### 08. Goroutines

8.1 - OS Scheduler Mechanics
8.2 - Go Scheduler Mechanics
8.3 - Creating Go Routines

### 09. Data Races

9.1 - Managing Data Races

### 10. Channels

10.1 - Signaling Semantics
10.2 - Basic Patterns
10.3 - Fan Out
10.4 - Wait for Task
10.5 - Pooling
10.6 - Fan Out Semaphore
10.7 - Fan Out Bounded
10.8 - Drop Pattern
10.9 - Cancellation Pattern

### 11. Concurrency Patterns

11.1 - Failure Detection

### 12. Testing

12.1 - Basic Unit Testing
12.2 - Table Unit Testing
12.3 - Mocking Web Server Response
12.4 - Testing Internal Endpoints
12.5 - Sub Tests
12.6 - Code Coverage

### 13. Benchmarking

13.1 - Basic Benchmarking
13.2 - Validate Benchmarking
13.3 - CPU-Bound Benchmarking
13.4 - IO-Bound Benchmarking

### 14. Profiling & Tracing

14.1 Profiling Guidelines
14.2 Stack Traces
14.3 Micro Level Optimization
14.4 Macro Level Optimization
14.5 Execution Tracing

## Ultimate Service with Kubernetes

This course teaches you how to build production-level services in Go leveraging the power of Kubernetes. See The instructor walking through the design philosophies and guidelines for building services in Go.

### 01. Introduction

1.0 - Intro
1.1: Design Philosophy, Guidelines, What to Expect
1.2: Tooling to Install

### 02. Modules

2.0 - Intro
2.1: Adding Dependencies
2.2: Module Mirrors
2.3: Checksum Database
2.4: Vendoring
2.5: MVS Algorithm

### 03. Kubernetes

3.0 - Intro
3.1: Tooling Installation
3.2: Understanding Clusters, Nodes and Pods
3.3: Write Basic Service for Testing
3.4: Zarf Layer
3.4.1: Docker Images
3.4.2: Kind Configuration
3.4.3: Core K8s Configuration
3.4.4: K8s Quotas / Patching

### 04. Initial Service Design

4.0 - Intro
4.1: Project Layers, Policies, and Guidelines
4.2: Prepare Project
4.3: Logging Support
4.4: Configuration Support
4.5: Debugging / Metrics Support
4.6: Shutdown Signaling and Load Shedding

### 05. HTTP Routing Basics

5.0 - Intro
5.1: Basic Structure of an HTTP Router
5.2: Add a Readiness, Liveness and Test Handler

### 06. Web Framework

6.0 - Intro
6.1: Custom Router
6.2: Custom Handler Function
6.3: Middleware Support
6.4: Sending Responses

### 07. Middleware

7.0 - Intro
7.1: Logging
7.2: Request Context
7.3: Error Handling
7.3.1: Understanding what Error Handling Means
7.3.2: Declaring Custom Error Types
7.3.3: Consistent Handling and Response
7.4: Panic Handling
7.5: Metrics

### 08. JSON Web Tokens

8.0 - Intro
8.1: Understanding JWT
8.2: Private/Public Key Generation
8.3: Token Generation
8.4: Token Signature Validation

### 09. Authentication / Authorization

9.0 - Intro
9.1: Auth Package
9.2: Implementation of an In-Memory Key Store
9.3: Middleware
9.4: Auth Unit Test

### 10. Database Support

10.0 - Intro
10.1: Kubernetes Support for Postgres
10.2: Using Sqlx
10.3: Update Readiness Handler to Perform DB Checks

### 11. Database Migrations and Seeding

11.0 - Intro
11.1: Maintaining Database Schemas
11.2: Seeding Data
11.3: Init Containers

### 12. Business Packages

11.0 - Intro
11.1: Maintaining Database Schemas
11.2: Seeding Data
11.3: Init Containers

### 13. Testing Data Business Packages

13.0 - Intro
13.1: Support for Starting and Stopping Containers
13.2: Support for Starting and Stopping a Unit Test
13.3: Write User CRUD Data Unit Tests

### 14. REST API

14.0 - Intro
14.1: Writing User Web Handlers
14.2: Support for Starting and Stopping an Integration Test
14.3: Write Integration Tests for Users

### 15. Open Telemetry

15.0 - Intro
15.1: Integrate OTEL Web Handler into the Framework
15.2: Integrate OTEL into Service Startup
15.3: Add Zipkin into POD
15.4: Add Tracing Calls Inside Functions to Trace

### 16. Review Services Project

16.0 - Intro
16.1: Check For Dependcy Upgrades
16.2: Rebuild and Run the Project

### 17. Beyond the Branch

17.0 - Intro
17.1: Introduction
17.2: Handling Log Sync on Error
17.3: Removing DB Driver Types
17.4: New POD for Zipkin
17.5: v1 for Business/Web
17.6: Data and Core Changes

## Writing Secure Go Code

For engineers who want to improve the security of theur Go applications.

### 01. Whats Covered

1.0 - The Security Mindset
2.0 - Go Security Policy
3.0 - OWASP Top Ten
4.0 - Input
5.0 - Output
6.0 - Authentication
7.0 - Infrastructure
