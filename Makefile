UNAME_S := $(shell uname -s | tr 'A-Z' 'a-z')

build: build-$(UNAME_S)

run:
	go run playstore-ratings.go

test:
	go test

all: build-windows build-linux build-darwin docs

clean:
	rm -Rf bin

build-%:
	@echo Building for $(*)
	@if [ "$(*)" == "windows" ]; then \
	echo GOOS=$(*) GOARCH=amd64 go build -o bin/playstore-ratings-$(*).exe; \
	GOOS=$(*) GOARCH=amd64 go build -o bin/playstore-ratings-$(*).exe; \
	else \
	echo GOOS=$* GOARCH=amd64 go build -o bin/playstore-ratings-$*$(SUFFIX); \
	GOOS=$* GOARCH=amd64 go build -o bin/playstore-ratings-$*$(SUFFIX); \
	fi

docs:
ifeq (, $(shell which doctoc))
 $(error "No doctoc in $(PATH), consider installing to run 'docs' target")
endif
	doctoc readme.md