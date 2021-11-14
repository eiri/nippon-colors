.DEFAULT_GOAL := run

PROJ := nipponcolors
SRCS := $(shell find $(CURDIR) -name '*.go')

$(PROJ): color.json $(SRCS)
	go build -o $@ ./cmd/$(PROJ)/...

.PHONY: run
run: $(PROJ)
	./$(PROJ) demo

.PHONY: clean
clean:
	go clean
	rm -f $(CURDIR)/$(PROJ)

color.js:
	curl -L https://raw.githubusercontent.com/ssshooter/nippon-color/master/src/data/color.js -o $@

color.json: color.js
	go run ./cmd/convert/main.go $< | jq -cM '.' > $@
