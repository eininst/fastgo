registry_mirrors = registry.cn-zhangjiakou.aliyuncs.com
username = youyin319
namespace = eininst
group = g

start:
	sh $(CURDIR)/build/start.sh ${env}

swagger:
	sh $(CURDIR)/scripts/swagger.sh

doc:
	swag init -g router.go -d ./api/${app} -o ./api/${app} -ot json

init:
	sh $(CURDIR)/scripts/swarm.sh

deploy:
	docker stack deploy --with-registry-auth -c deployments/$(yml).yml $(group)

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
	$(group)_${app}

stop:
ifeq (${srv},)
	docker stack rm $(group)
else
	docker service rm $(group)_${srv}
endif

fab:
	make build app=${app} f=${f}
	make push app=${app}
	make update yml=${yml} app=${app}

log:
	docker service logs -f --tail 200 $(group)_${app}

login:
	docker login --username=$(username) $(registry_mirrors)

clean:
	yes | docker system prune

.PHONY: swagger doc init build push deploy update stop fab log login clean