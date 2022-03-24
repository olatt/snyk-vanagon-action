containerName = sec-van-action
# testFolder = /Users/jeremy.mill/Documents/bolt-vanagon/
# testFolder = /Users/jeremy.mill/Documents/puppet-runtime/
# testFolder = /Users/jeremy.mill/Documents/pe-installer-vanagon/
# testFolder = /Users/jeremy.mill/Documents/pxp-agent-vanagon/
testFolder = /Users/oak.latt/dev/pe-opsworks-tools-vanagon/

SSHKEY := $(shell cat /Users/oak.latt/.ssh/id_ed25519 | base64)
clean:
	-rm vanagon_action
	-rm header_proxy/header_proxy
	-rm -rf ./testfiles/repo
	-docker rm $(containerName)
build:
	make clean
	env GOOS=linux go build -o vanagon_action
	# cd header_proxy; env GOOS=linux go build -o header_proxy
	docker build --platform linux/amd64 -t $(containerName) .
copy_testfiles:
	-mkdir -p ./testfiles/repo
	cp -r $(testFolder) ./testfiles/repo
itest:
	make clean
	make build
	make copy_testfiles
	docker run --platform linux/amd64 -i --name $(containerName) \
		-e INPUT_SNYKORG=sectest \
		-e INPUT_SNYKTOKEN=$(SNYK_TOKEN) \
		-e GITHUB_WORKSPACE=/github/workspace \
		-e INPUT_SKIPPROJECTS=agent-runtime-1.10.x,agent-runtime-5.5.x,bolt-runtime,client-tools-runtime-2019.8.x,client-tools-runtime-irving,client-tools-runtime-main,pdk-runtime,pe-bolt-server-runtime-2019.8.x,pe-bolt-server-runtime-main,pe-installer-runtime-2019.8.x,pe-installer-runtime-main,agent-runtime-main \
		-e INPUT_SKIPPLATFORMS=cisco-wrlinux-5-x86_64,cisco-wrlinux-7-x86_64,debian-10-armhf,eos-4-i386,fedora-30-x86_64,fedora-31-x86_64,osx-10.14-x86_64 \
		-e INPUT_SSHKEY="$(SSHKEY)" \
		-e INPUT_SSHKEYNAME=id_ed25519 \
		-e INPUT_SVDEBUG=true \
		-e INPUT_BRANCH=main-test \
		-v "/Users/oak.latt/dev/security-snyk-vanagon-action/testfiles/repo":"/github/workspace" \
		-t $(containerName) 

exec:
	make clean
	make build
	make copy_testfiles
	docker run --platform linux/amd64 --name $(containerName) \
		-e INPUT_SNYKORG=sectest \
		-e INPUT_SNYKTOKEN=$(SNYK_TOKEN) \
		-e GITHUB_WORKSPACE=/github/workspace \
		-e INPUT_SKIPPROJECTS=agent-runtime-1.10.x,agent-runtime-5.5.x,bolt-runtime,client-tools-runtime-2019.8.x,client-tools-runtime-irving,client-tools-runtime-main,pdk-runtime,pe-bolt-server-runtime-2019.8.x,pe-bolt-server-runtime-main,pe-installer-runtime-2019.8.x,pe-installer-runtime-main,agent-runtime-main \
		-e INPUT_SKIPPLATFORMS=cisco-wrlinux-5-x86_64,cisco-wrlinux-7-x86_64,debian-10-armhf,eos-4-i386,fedora-30-x86_64,fedora-31-x86_64,osx-10.14-x86_64 \
		-e INPUT_SSHKEY="$(SSHKEY)" \
		-e INPUT_SSHKEYNAME=id_ed25519 \
		-e INPUT_SVDEBUG=true \
		-e INPUT_BRANCH=main \
		-v "/Users/oak.latt/dev/security-snyk-vanagon-action/testfiles/repo":"/github/workspace" \
		-it $(containerName) /bin/bash

test:
	echo $(SSHKEY)

start_cache:
	docker rm cache_proxy
	docker build -t cache_proxy -f ./dev_tools/Dockerfile .
	docker run -it --rm --name cache_proxy -e INPUT_RPROXYUSER=artifactory -e INPUT_RPROXYKEY="$(RPROXY_KEY)" -p8080:80 cache_proxy
