# In-memory cache with HTTP interface
#### Overview
- In-memory cache with HTTP interface written in Go
- Entries in cache are evicted based on a configurable TTL
- HTTP GET /{key} replies with the value as body or 404 if no such key exists
- HTTP POST /{key} accepts and stores UTF-8 body (text/plain;charset=utf-8) for the key

Uses the following frameworks
- Configuration: Viper (https://github.com/spf13/viper)
- Caching: Bigcache (https://github.com/allegro/bigcache)
- HTTP routing: gorilla/mux (https://github.com/gorilla/mux) 

#### Environment variables

The following environments are supported:
- Development. Set ENV=DEV before launching.

#### Run tests
`export ENV=DEV` <br> 
`go test ./... -v`

#### Launch app
`export ENV=DEV` <br> 
`go run main.go`

#### Config
See config.development.yml

#### Example usage

POST <br>
`curl -v -d '{"name":"value1", "description":"value2"}' -H "Content-Type: text/plain;charset=utf-8" -X POST http://localhost:8080/my-key` <br><br>
GET <br>
`curl -v http://localhost:8080/my-key`

#### Design considerations
Bigcache is used as implementation for the key value store. It has several benefits over using a ordinary go map when it comes to concurrent access and limiting expensive GC cycles. <br>
For more info see this blogpost: https://dev.to/douglasmakey/how-bigcache-avoids-expensive-gc-cycles-and-speeds-up-concurrent-access-in-go-12bb
 
#### Limitations
- Https not supported
- Authentication not implemented (yet :) )