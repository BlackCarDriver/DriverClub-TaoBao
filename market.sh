#!/bin/bash

# It shell is used to run the taobao project
# argument can choise : 
# stop	: stop the running market-server container
# build : build go code
# run	: run go server
# ngbd	: npm build for angular project
# ngit	: npm install for angular project

isngit=0
isngbd=0
isbuild=0
isstop=0
isrun=0

for var in "$@"
do
t="$var"

# npm install for angular project
if [ ${t} == "ngit" ] && [ ${isngit} -eq 0 ]
then 
	isngit=1
	echo "going to npm install=>"
	sudo docker run -it --rm \
	-v $PWD/angular:/workplace \
    blackcardriver2/mcis:ngcli \
	sh -c 'cd /workplace && npm install --verbose'
	if [ $? -ne 0 ]
	then
		echo "npm install fail!"
		exit $?
	else
		echo "npm install success!"
	fi
	continue
fi

# build angular project
if [ ${t} == "ngbd" ] && [ ${isngbd} -eq 0 ]
then 
	isngbd=1
	echo "going to build Angular project=>"
	sudo docker run -it --rm \
	-v $PWD/angular:/workplace \
    blackcardriver2/mcis:ngcli \
	sh -c 'cd /workplace && npm run-script build'
	if [ $? -ne 0 ]
	then
		echo "Angular project build fail!"
		exit $?
	else
		echo "Angular build success!"
	fi
	continue
fi

# stop the old server continer
if [ ${t} == "stop" ] && [ ${isstop} -eq 0 ]
then
	isstop=1
    echo "going to kill old container=>"
	tmpval=$( sudo docker ps | grep 'market-server')
	if [ ${#tmpval} != 0 ]
	then 
		sudo docker stop market-server
		sleep 3s 
		if [ $? -ne 0 ]
		then 
			echo 'kill old server success'
		else
			echo "kill old server fail"
			exit $?
		fi
	else
		echo "old server not found"
	fi
	unset tmpval
	continue
fi

# build server
if [ ${t} == "build" ] && [ ${isbuild} -eq 0 ]
then
	isbuild=1
	echo "go server is going to build =>"
	docker run -it \
	--rm --name gobuild \
	-v $PWD/Go/src:/go/src \
	golang:alpine \
	sh -c 'cd /go/src && go build TaobaoServer/main.go'
	if [ $? -ne 0 ]
	then 
		echo 'go build success!'
	else
		echo "go build fail"
		exit $?
	fi
	continue
fi

# run go server
if [ ${t} == "run" ] && [ ${isrun} -eq 0 ]
then
	isrun=0
	echo "go server is going to run =>"
	docker run \
	--rm \
	--name market-server \
	-v /home/ubuntu/DockerWorkPlace/Market/DriverClub-taobao/Go/src:/workplace \
	-v /home/ubuntu/Nginx/source/marketUpload:/source \
	--network market-net \
	-d -p 4749:4747 \
	alpine:latest \
	sh -c 'mv /workplace/main /workplace/TaobaoServer/main && cd /workplace/TaobaoServer && ./main'
	if [ $? -ne 0 ]
	then 
		echo 'docker run success!'
	else
		echo "docker run fail"
		exit $?
	fi
	continue
else
	echo "Unknow flag: ${t}"
fi

done