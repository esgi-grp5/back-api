# back-api
Golang

# need to do that to use swagger
export PATH=$(go env GOPATH)/bin:$PATH
swag init -g cmd/games/main.go --output swagger/games
swag init -g cmd/user/main.go --output swagger/user