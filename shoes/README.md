[comment]: <> (This is a generated file please edit source in ./templates)
[comment]: <> (All modification will be lost, you have been warned)
[comment]: <> ()
### Sample CRUD API for the mysql database doadmin:AVNS_nQrjtn8ilVHqYs6xIim@tcp(dbaas-db-7154856-do-user-13260059-0.b.db.ondigitalocean.com:25060)/shoes_development?parseTime=true

## Example
The project is a RESTful api for accessing the mysql database doadmin:AVNS_nQrjtn8ilVHqYs6xIim@tcp(dbaas-db-7154856-do-user-13260059-0.b.db.ondigitalocean.com:25060)/shoes_development?parseTime=true.

## Project Files
The generated project will contain the following code under the `./example` directory.
* Makefile
  * useful Makefile for installing tools building project etc. Issue `make` to display help
* .gitignore
  * git ignore for go project
* go.mod
  * go module setup, pass `--module` flag for setting the project module default `example.com/example`
* README.md
  * Project readme
* app/server/main.go
  * Sample Gin Server, with swagger init and comments
* api/*.go
  * REST crud controllers
* dao/*.go
  * DAO functions providing CRUD access to database
* model/*.go
  * Structs representing a row for each database table

The REST api server utilizes the Gin framework, GORM db api and Swag for providing swagger documentation
* [Gin](https://github.com/gin-gonic/gin)
* [Swaggo](https://github.com/swaggo/swag)
* [Gorm](https://github.com/jinzhu/gorm)

## Building
```.bash
make example
```
Will create a binary `./bin/example`

## Running
```.bash
./bin/example
```
This will launch the web server on xinqi.dev:443

## Swagger
The swagger web ui contains the documentation for the http server, it also provides an interactive interface to exercise the api and view results.
https://xinqi.dev:443/swagger/index.html

## REST urls for fetching data


* https://xinqi.dev:443/arinternalmetadata_
* https://xinqi.dev:443/schemamigrations_
* https://xinqi.dev:443/shoes

## Project Generated Details
```.bash
gen \
    --sqltype=mysql \
    --connstr=doadmin:AVNS_nQrjtn8ilVHqYs6xIim@tcp(dbaas-db-7154856-do-user-13260059-0.b.db.ondigitalocean.com:25060)/shoes_development?parseTime=true \
    --database=shoes_development \
    --templateDir=./templates \
    --model=model \
    --dao=dao \
    --api=api \
    --out=./ \
    --module=shoes \
    --json \
    --json-fmt=snake \
    --gorm \
    --guregu \
    --mod \
    --makefile \
    --server \
    --overwrite \
    --host=xinqi.dev \
    --port=443 \
    --rest \
    --listen=:8080 \
    --scheme=https \
    --generate-dao \
    --generate-proj \
    --file_naming={{.}} \
    --model_naming={{FmtFieldName .}} \
    --swagger_version=1.0 \
    --swagger_path=/ \
    --swagger_tos= \
    --swagger_contact_name=Me \
    --swagger_contact_url=http://me.com/terms.html \
    --swagger_contact_email=me@me.com
```











