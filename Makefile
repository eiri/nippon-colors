.DEFAULT_GOAL := demo

PROJ := nipponcolors
SRCS := $(shell find $(CURDIR) -name '*.go')

$(PROJ): color.json const.go $(SRCS)
	go build -o $@ ./cmd/$(PROJ)/...

.PHONY: demo
demo: $(PROJ)
	./$(PROJ) demo

.PHONY: clean
clean:
	go clean
	rm -f $(CURDIR)/$(PROJ)

color.js:
	curl -L https://raw.githubusercontent.com/ssshooter/nippon-color/master/src/data/color.js -o $@

color.json: color.js convert
	./convert to_json -i $< -o $@

const.go: color.json convert
	./convert to_const -i $< -o $@

convert:
	go build -o $@ ./cmd/convert/...