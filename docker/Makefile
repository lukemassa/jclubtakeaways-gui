include Makefile.config

.PHONY: build push

build:
	docker build -t $(NS)/$(REPO):$(VERSION) .

push:
	docker push $(NS)/$(REPO):$(VERSION)

run:
	docker run --rm --name $(NAME)-shell $(VOLUMES) $(NS)/$(REPO):$(VERSION) ./bin/report_who_can ndmad2 zookeeper-operator 

email:
	docker run --rm --name $(NAME)-shell $(VOLUMES) $(NS)/$(REPO):$(VERSION) ./bin/report_who_can ndmad2 zookeeper-operator --email clin@tripadvisor.com

shell:
	docker run --rm --name $(NAME)-shell $(VOLUMES) -it $(NS)/$(REPO):$(VERSION) sh
