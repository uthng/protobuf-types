PROTOC_OPTS ?= -I $(PWD)
PROTO_DIR = $(shell pwd)

proto-gen:
	PROTO_FILES=`find proto -name "*.proto"`; \
	for pb in $$PROTO_FILES; do \
	    protoc ${PROTOC_OPTS} --go_out=$(PROTO_DIR) $(PROTO_DIR)/$$pb; \
	done; \
