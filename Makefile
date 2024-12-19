html:
	@go run main.go

development:
	git rev-parse HEAD > $(CURDIR)/docs/commit.txt
