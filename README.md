## Background

Repository to work on energy API found here https://consumerdatastandardsaustralia.github.io/standards/#energy-apis

NOTE: Use the `.json` version of the OpenAPI spec and not the `.yaml`

## Building

To regenerate the OpenAPI code remove the `pkg/` directory

### Docker

```
docker run   --rm \
-v ${PWD}:/local  --user $(id -u):$(id -g)  openapitools/openapi-generator-cli generate \
-i /local/cdsenergy.json  \
-g go-server \
-o /local/pkg \
-p serverPort=8888 -p packageName=energyserver  -p router=chi -p onlyInterfaces=false \
-p sourceFolder=pkg --git-repo-id cdsenergy/v1 --git-user-id nanikjava
```

```
rm pkg/.openapi-generator-ignore pkg/Dockerfile  pkg/go.* pkg/main.go pkg/README.md  && rm pkg/.openapi-generator -rf  && rm pkg/api -rf
```
### go-server

When using Docker to run openapi-generator if the generated source code are copied as `root` into local drive then compiled the code locally
and run it using `java -jar`

```
java -jar <open-api-generator-location>/target/openapi-generator-cli.jar generate   \ 
 -i ./cdsenergy.json     -g go-server  -o $PWD/pkg  \
 -p serverPort=8888 -p packageName=energyserver  -p router=chi -p onlyInterfaces=false \
 -p sourceFolder=pkg --git-repo-id cdsenergy/v1 --git-user-id nanikjava 
```

After generation the code remove the following files from `pkg/` folder

```
rm pkg/.openapi-generator-ignore pkg/Dockerfile  pkg/go.mod pkg/go.sum pkg/main.go pkg/README.md pkg/api/openapi.yaml && rm pkg/.openapi-generator -rf 
```

### Issues

Following issues found when running openapi-generator for Go

* Any spec that has `tags` containing multiple values like this `"tags" : [ "Energy", "Generic Tariffs" ]` will create duplicate route and functions. To resolve the `tags` must be
  `"tags" : [ { "name" :  "Energy" }, { "name" :  "Generic Tariffs"} ]`
* Inlined enum type are not generated so we need to create our own type. This does not work
```
   "schema" : {
     "default" : "ALL",
     "enum" : [ "STANDING", "MARKET", "REGULATED", "ALL" ],
     "type" : "string"
   }
```

need to be modify to

```
  "schema" : {
     "default" : "ALL",
     "$ref" : "#/components/schemas/ListPlanType",
     "type" : "object"
   }
```

and need to define the `ListPlanType` in `schemas` as follows


```
 "schemas" : {
   "ListPlanType" : {
       "enum" : [ "STANDING", "MARKET", "REGULATED", "ALL_LISTPLAN" ],
       "type" : "string"
   },
   ...
  }        
```