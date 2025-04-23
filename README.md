# OpenAPI/gRPC
## Table of content
- [OpenAPI](#openapi)
- [gRPC](#grpc)

## OpenAPI
OpenAPI Specification (formerly Swagger Specification) is an API description format for REST APIs. An OpenAPI file allows you to describe your entire API, including:
- Available endpoints and operations on each endpoint
- Operation parameters Input and output for each operation
- Authentication methods
- Contact information, license, terms of use, and other information.
API specifications can be written in YAML or JSON.\
The OpenAPI Specification (OAS) allows the description of a remote API accessible through HTTP or HTTP-like protocols. This description, which may be stored as one or more documents (such as local files or HTTP-accessible network resources), is called an OpenAPI Description (OAD). Every OpenAPI Descriptions must contain an OpenAPI Object with at least the fields `openapi`, and `info`, and either `paths`, `components` or `webhooks`.

## gRPC
In gRPC, a client application can directly call a method on a server application on a different machine. On the server side, the server implements this interface and runs a gRPC server to handle client calls. On the client side, the client has a stub (referred to as just a client in some languages) that provides the same methods as the server.\
As in many RPC systems, gRPC is based around the idea of defining a service, specifying the methods that can be called remotely with their parameters and return types. By default, gRPC uses Protocol Buffers, Google’s mature open source mechanism for serializing structured data.\
gRPC can be defined 4 kinds of service method: Unary RPC, Server streaming RPC, Client streaming RPC, Bidirectional streaming RPC
### Unary RPC
First consider the simplest type of RPC where the client sends a single request and gets back a single response.
- Once the client calls a stub method, the server is notified that the RPC has been invoked with the client’s metadata for this call, the method name, and the specified deadline if applicable.
- The server can then either send back its own initial metadata straight away, or wait for the client’s request message.
- Once the server has the client’s request message, it creates and populates a response. The response is then returned to the client together with status details (status code and optional status message) and optional trailing metadata.
- If the response status is OK, then the client gets the response, which completes the call on the client side.

### Server streaming RPC
A server-streaming RPC is similar to a unary RPC, except that the server returns a stream of messages in response to a client’s request. After sending all its messages, the server’s status details and optional trailing metadata are sent to the client.

### Client streaming RPC
A client-streaming RPC is similar to a unary RPC, except that the client sends a stream of messages to the server instead of a single message. The server responds with a single message after it has received all the client’s messages.

### Bidirectional streaming RPC
In a bidirectional streaming RPC, the call is initiated by the client invoking the method and the server receiving the client metadata, method name, and deadline. The server can choose to send back its initial metadata or wait for the client to start streaming messages. Client- and server-side stream processing is application specific. Since the two streams are independent, the client and server can read and write messages in any order.

### RPC Cancellation
Either the client or the server can cancel an RPC at any time. A cancellation terminates the RPC immediately so that no further work is done. Changes made before a cancellation are not rolled back.

### Metadata
Metadata is information about a particular RPC call in the form of a list of key-value pairs, where the keys are strings and the values are typically strings, but can be binary data. Keys are case insensitive and consist of ASCII letters, digits, and special characters `-`, `_`, `.` and must not start with `grpc-` (which is reserved for gRPC itself). Binary-valued keys end in `-bin` while ASCII-valued keys do not. User-defined metadata is not used by gRPC, which allows the client to provide information associated with the call to the server and vice versa.

### Channel
A gRPC channel provides a connection to a gRPC server on a specified host and port. It is used when creating a client stub. Clients can specify channel arguments to modify gRPC’s default behavior, such as switching message compression on or off. A channel has state, including `connected` and `idle`.