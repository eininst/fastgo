export APP=$1

make build app=nginx f=nginx/Dockerfile
make build app=helloword

make push app=nginx
make push app=helloword

make deploy yml=nginx
make deploy yml=api

docker service ls


make fab app=helloword yml=api
