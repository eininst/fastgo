export ENV=$1
#make login
#make init

make build app=nginx f=nginx/Dockerfile
make build app=helloword

make push app=nginx
make push app=helloword

make deploy yml=nginx
make deploy yml=api

docker service ls
