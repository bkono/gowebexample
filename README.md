# gowebexample

Super simple example of the differences between `micro web` and `micro api --handler http` from the
service perspective

When running `micro web` and calling the service, the url will be handled from a func matching on
`/upload`

```
❯ curl localhost:8082/image/upload
hello from upload, the url is "/upload"
```

When running `micro api --handler http --namespace go.micro.web` and calling the service, the url
will be handled from a func matching on `/image/upload`, note the included service name. This is due
to the `--handler http` functioning as a reverse proxy, rather than using the service part of the
path for discovery/routing and then removing it before continuing the call.

```
❯ curl localhost:8080/image/upload
hello from upload, the url is "/image/upload"
```
