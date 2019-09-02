# ng build 
docker run \
--rm --name ngbuild \
-v /home/ubuntu/DockerWorkPlace/jenkins/jenkins_home/workspace/market/angular:/workplace \
-w /workplace \
blackcardriver2/mcis:ngcli \
sh -c 'ls && pwd && ng build --prod --aot --base-href /market/'

if [ $? -ne 0 ]; then exit 1; fi

# copy build result
if [ -w /nginx_home/market ]; then rm -rf /nginx_home/market && echo 'rm market old dist succeed!'; fi
cp -rf angular/dist/mywebsite /nginx_home/market

# go build 
docker run \
--rm --name gobuild \
-v /home/ubuntu/DockerWorkPlace/jenkins/jenkins_home/workspace/market/Go/src:/go/src \
golang:alpine \
sh -c 'cd /go/src && go build TaobaoServer/main.go'

if [ $? -ne 0 ]; then exit 2; fi

# go run 
docker run \
--rm \
--name market-setver \
-v /home/ubuntu/DockerWorkPlace/jenkins/jenkins_home/workspace/market/Go/src:/workplace \
-p 4749:4747 \
alpine:latest \
sh -c 'cp /workplace/main /workplace/TaobaoServer/main && cd /workplace/TaobaoServer && ./main'
