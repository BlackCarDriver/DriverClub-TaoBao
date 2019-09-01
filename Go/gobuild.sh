sudo docker run \
--rm -v \
$PWD/src:/go/src \
golang:alpine \
sh -c 'cd /go/src && go build TaobaoServer/main.go'
