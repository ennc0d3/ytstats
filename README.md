# Youtube stat tool

Not just another stat tool. The idea is to do some fun with energy consumption.


## Generating client and stub code from OpenApi
```
https://generator.swagger.io/#/
```

### Generate go-server stub
```
 curl -X POST -H "content-type:application/json" -d '{"swaggerUrl":"https://petstore.swagger.io/v2/swagger.json"}' https://generator.swagger.io/api/gen/servers/go-server

```

## Start a local DB instance
```
docker run --name petsitter-db -d -p 27017:27017 mongo:latest

```
