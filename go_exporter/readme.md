# Steps

* Source
[Prometheus instrumentation](https://prometheus.io/docs/guides/go-application/#instrumenting-a-go-application-for-prometheus)

* Create `go.mod`
```bash
go mod init dejanualex/goexporter
```
* Install dependencies
```bash
# install gin, previously declared in the import
go get .
# explicitly get the dependecy using semantic version
go get github.com/gin-gonic/gin@v1.7.2
```