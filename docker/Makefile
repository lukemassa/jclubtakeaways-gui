include Makefile.config

.PHONY: build push

build:
	docker build -t $(NS)/$(REPO):$(VERSION) .

push:
	docker push $(NS)/$(REPO):$(VERSION)

shell:
	docker run --rm --name $(NAME)-shell $(VOLUMES) -it $(NS)/$(REPO):$(VERSION) sh
