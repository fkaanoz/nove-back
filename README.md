## TODO
- Reverse the middleware order.
- Redis Client should be implemented.



- Errors from API should be handled in Frontend with toasters or something like that.  - 10th Nov.

### Error in Build Stage (for darwin):
- confluent kafka package needs CGO_ENABLED for build process. However golang image for Darwin encounters a problem related with gcc. When you build the eventual image in an Ubuntu Server, problem disappear. 