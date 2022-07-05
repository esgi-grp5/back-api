# back-api
Golang

# need to do that to use swagger
export PATH=$(go env GOPATH)/bin:$PATH
swag init -g server/main.go --output swagger/users