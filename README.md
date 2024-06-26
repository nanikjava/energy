## Background

Repository to work on energy API found here https://consumerdatastandardsaustralia.github.io/standards/#energy-apis

NOTE: Use the `.json` version of the OpenAPI spec and not the `.yaml`

## Building

### go-server

When using Docker to run openapi-generator the generated source code are copied as `root` into local drive. Due to this issue decided to just compiled the code locally
and run it using `java -jar`

```
java -jar /home/nanik/GolandProjects/openapi-generator/modules/openapi-generator-cli/target/openapi-generator-cli.jar generate   \ 
 -i ./cdsenergy.json     -g go-server  -o $PWD/pkg  \
 -p serverPort=8888 -p packageName=energyserver  -p router=chi -p onlyInterfaces=false \
 -p sourceFolder=pkg --git-repo-id cdsenergy/v1 --git-user-id nanikjava  \
```

After generation the code remove the following files from `pkg/` folder

```
pkg/.openapi-generator-ignore
pkg/Dockerfile
pkg/go.mod
pkg/main.go
pkg/README.md
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