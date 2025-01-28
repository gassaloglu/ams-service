@echo off
echo Running 'go get' to install gin package...
go get -u github.com/gin-gonic/gin
go get go.mongodb.org/mongo-driver/mongo
go get firebase.google.com/go/v4
go get github.com/spf13/viper
pause
