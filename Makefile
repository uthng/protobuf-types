PROTOC_OPTS ?= -I $(PWD)
PROTO_DIR = $(shell pwd)

proto-gen:
	PROTO_FILES=`find proto -name "*.proto"`; \
	for pb in $$PROTO_FILES; do \
	    protoc ${PROTOC_OPTS} --go_out=$(GOSRC) $(PROTO_DIR)/$$pb; \
	done; \

example:
	make -C example all

.PHONY : example proto-gen
