## Background

Repository to work on energy API found here https://consumerdatastandardsaustralia.github.io/standards/#energy-apis

## Building

### go-gin-server

```
openapi-generator-cli generate     -i ./cdsenergy.yaml     -g go-gin-server     \
    -o $PWD/pkg  -p serverPort=8888 -p packageName=energyserver  \
    -p apiPath=gen -p interfaceOnly=true \
    --git-repo-id cdsenergy/v1 --git-user-id nanikjava 
```
