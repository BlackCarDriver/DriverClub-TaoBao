# ngbuild
docker run -it \
--rm --name ngbuild \
-v $PWD/angular:/workplace \
-w /workplace \
blackcardriver2/mcis:ngcli \
sh -c 'cd /workplace &&  ng build --prod --aot --base-href /market/'

#if [ $? -ne 0 ]; then echo 'ng build fail!' exit 1; fi

# move dist file to nginx
sudo rm -rf  /home/ubuntu/Nginx/html/market
sudo mv $PWD/angular/dist/market /home/ubuntu/Nginx/html/market

# go build
docker run -it \
--rm --name gobuild \
-v /home/ubuntu/DockerWorkPlace/Market/DriverClub-taobao/Go/src:/go/src \
golang:alpine \
sh -c 'cd /go/src && go build TaobaoServer/main.go'

if [ $? -ne 0 ]; then echo 'Gobuild fail!' exit 2; fi

# stop old container
tmpval=$( sudo docker ps | grep 'market-server')
if [  ${#tmpval} != 0  ]; then docker stop market-server && sleep 3s && echo 'kill rinning market-server' ; fi

# run server
docker run -d \
--rm \
--name market-server \
-v /home/ubuntu/DockerWorkPlace/Market/DriverClub-taobao/Go/src:/workplace \
-p 4749:4747 \
alpine:latest \
sh -c 'cp /workplace/main /workplace/TaobaoServer/main && cd /workplace/TaobaoServer && ./main'
