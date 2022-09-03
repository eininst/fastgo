registry_mirrors = registry.cn-zhangjiakou.aliyuncs.com
username = youyin319
namespace = eininst

start:
	sh $(CURDIR)/build/start.sh ${env}

swagger:
	sh $(CURDIR)/scripts/swagger.sh

doc:
	swag init -g router.go -d ./api/helloword -o ./docs -ot json --instanceName helloword

init:
	sh $(CURDIR)/scripts/swarm.sh

deploy:
	docker stack deploy --with-registry-auth -c deployments/$(yml).yml srv

build:
ifeq (${f},)
	docker build -f Dockerfile --build-arg APP=${app} -t ${app} .
else
	docker build -f ${f} --build-arg APP=${app} -t ${app} .
endif


push:
	docker tag ${app} $(registry_mirrors)/$(namespace)/${app}:latest
	docker push $(registry_mirrors)/$(namespace)/${app}:latest

update:
	docker service update --force --image \
	`sh $(CURDIR)/scripts/yaml.sh $(CURDIR)/deployments/${yml}.yml services_$(app)_image` \
	srv_${app}

stop:
ifeq (${srv},)
	docker stack rm srv
else
	docker service rm srv_${srv}
endif

fab:
	make build app=${app}
	make push app=${app}
	make update yml=${yml} app=${app}

log:
	docker service logs -f --tail 200 srv_${app}

login:
	docker login --username=$(username) $(registry_mirrors)

clean:
	yes | docker system prune

.PHONY: swagger doc init build push deploy update stop fab log login clean