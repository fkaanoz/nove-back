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

## Other Notes:
- Repository will be refactored. It is not its eventual state.  
- pem files are added intentionally.
- **Commits are not too much tidy.** They are done when both frontend and backend are deployed to Azure.
  - It has a basic frontend which is written in Next.js to simulate project's distributed nature.
  - **Static pages are served with Next.js app and api requests are directed to Go Service by simple nginx configuration.**
- **Repository doesn't have too much comments.** 

### Error in Build Stage (for darwin):
- confluent kafka package needs CGO_ENABLED for build process. However golang image for Darwin encounters a problem related with gcc. When you build the eventual image in an Ubuntu Server, problem disappear.


## Cloc
**Note** : Vendor directory excluded.

Language|files|blank|comment|code
:-------|-------:|-------:|-------:|-------:
Go|46|345|48|1470
XML|7|0|0|144
Markdown|1|4|0|39
Bourne Shell|3|9|0|12
make|1|8|0|11
Dockerfile|1|4|0|7
--------|--------|--------|--------|--------
SUM:|59|370|48|1683


cloc|github.com/AlDanial/cloc v 1.96  T=0.02 s (2681.1 files/s, 95473.8 lines/s)
--- | ---





https://drive.google.com/file/d/1MTmTS3ts1s-ZEInLErwrkndZZsGox1tp/view?usp=sharing




