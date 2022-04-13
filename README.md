# go-api-base

Base code to run a RESTful or GRPC microservice application.

By combining this library and [go-api-frame](https://github.com/keenfury/go-api-frame) you can build REST/GRPC *go* microservice applications.  If you are familiar with RoR scaffolding, well this is the same concept but with *go*.

### Installation
```
git clone github.com/keenfury/go-api-base
```

### Usage
By design this code base is to be clone for each new service.

To clone the project:
```
cd scripts
./clone_this.sh github.com/keenfury/my_repo
```

One argument is required:
- full path from `$GOPATH/src` e.g. github.com/{git name}/my_repo

This will copy all files rename all the reference from `go-api-frame` to your new project name.

`my_repo` will be considered the *service name* from through the rest of these documents.

Open up your favorite development UI and load up your new project.  At this point the code can compile though it will do nothing.

This is where `go-api-frame` comes into play, follow the README at [go-api-frame](https://github.com/keenfury/go-api-frame). Most the 'heavy lifting' is done from this library.

Though you have full control of your code and change things in how best for your application, the `go-api-frame` and this library gets you going with the basic boilerplate code to do CRUD operations on a table or more if you would like.

The code has 'hook' lines that are comments (so don't delete them), that if so desired, you can rerun `go-api-frame` as many times as you want to add more groups of endpoints (as long as the table name is different, obviously).  The code will add the appropriate code to the following files:

- `cmd/rest/main.go`
- `cmd/grpc/main.go`
- `internal/v1/common.go`
- `internal/v1/<table name>.go` folder
- `pkg/proto/<service name>.proto`

If desired to remove a group of endpoints (CRUD operation for a table), these are the files you would need to alter.

Running this command:
```protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative %s.proto```

After the `<service name>.proto` file has changed, will update the `<service name>.grpc.pd.go` and `<service name>.pd.go` files as needed for the *grpc* service, to keep that generated code up-to-date.

### Engines
HTTP server using [Echo](https://echo.labstack.com/) the git repo is here: [git](https://github.com/labstack/echo). Echo's [benchmarking](https://github.com/vishr/web-framework-benchmark)

GRPC uses the GRPC V3 specification, see: [protobuf](https://developers.google.com/protocol-buffers/docs/reference/go-generated) and [grpc](https://grpc.io/docs/languages/go/basics/)

### Migration
This project also include a migration mechanics to do simple changes to your database each time your application starts up.  See `tools/migration/main.go` (shows up after you clone), see the notes in that file for more usage.  If you are familiar with RoR migration through active record it is similar, other than you write native SQL.  You can turn this off with the variable `UseMigration` in `config/config.go`.