all: ch

.PHONY: clean

ch:
	@sh -c "'$(CURDIR)/scripts/build.sh'"

clean:
	@rm -fv ./ch
