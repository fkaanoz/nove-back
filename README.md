## General Info
- It is a **backbone** service. It does **almost nothing**. 
- Repository follows **[William Kennedy](https://github.com/ardanlabs/service/)**'s folder structure.
- Entry point for service is ./app/shtil/main.go  - (completely irrelevant info : shtil is a Russian Song by Aria)
- Confluent Go Kafka library is used to connect multi-node kafka server.
  - There is a wrapper around library.
  - It does nothing meaningful with kafka. It just connects to server and when message comes it outputs message to standard output.
- Also, there is a redis wrapper around official redis library. 
  - It does nothing also. 
  - Some business related packages have redis connection. It might be used for caching costly database queries in these packages.
- Debug server uses standard library endpoints. 
  - One custom endpoint is added, it might be extended.
- app.go file in ./app/shtil/handlers adds handlers to mux. And api versioning is in there.
- business folder contains service specific packages.
- foundation folder contains more general and reusable packages.
- zarf folder contains deployment related things.


Some Packages:
- ardanlabs/conf for environment variables and configurations.
- sqlx for connecting to postgres.
- httptreemux as http multiplexer. 
  - It is a bare minimum http multiplexer, it does not have middleware concept. 
    - In foundation folder, it's ContextMux is embedded in a struct. And middlewares appended in its overridden Handle function.
    - In business folder, real middleware implementations are added.
  - It doesn't have a custom context other than standard library's request context.
    - In foundation/web/web.go, a custom context is added to hold tracking id and other stuffs.
- google/uuid for creating unique identifier (tracking id) for each request. 
- golang-jwt/jwt for stateful auth tokens. 
- uber.org/zap for structured logging.
- confluentinc/confluent-kafka-go for connecting to kafka.
- redis/go-redis for connecting to redis server.

## Notes:
- Repository will be refactored. It is not its eventual state.  
- pem files are added intentionally. 

### Error in Build Stage (for darwin):
- confluent kafka package needs CGO_ENABLED for build process. However golang image for Darwin encounters a problem related with gcc. When you build the eventual image in an Ubuntu Server, problem disappear.