PLUGINS=day01 day02 day03 day04

PLUGIN_LIBS=$(foreach p, $(PLUGINS), plugins/$(p).so)

SRCS = *.go plugin/*.go scanner/*.go

# .PHONY: all
# all: aoc2020

aoc2020: $(SRCS) $(PLUGIN_LIBS)
	go build -o $@

define plugin_template
plugins/$(1).so: $(1)/*.go $(SRCS)
	go build -buildmode=plugin -o plugins/ ./$(1)
endef

$(foreach p,$(PLUGINS), $(eval $(call plugin_template,$(p))))

.PHONY: clean
clean:
	rm -f $(PLUGIN_LIBS) aoc2020
