all: go rust

.PHONY: go rust

go:
	cd go && $(MAKE) bench

rust:
	cd rust && cargo bench
