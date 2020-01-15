.PHONY: all clean

PROTO_SRC_DIR=proto
PROTO_DST_DIR=proto

WARRIOR_PROG=warrior
WARRIOR_SRCS=warrior.go

all: gen-proto build_warrior

gen-proto:
	@echo "generating protobuf GO codes..."
	@protoc -I=$(PROTO_SRC_DIR) --go_out=$(PROTO_DST_DIR) $(PROTO_SRC_DIR)/*.proto

build_warrior:
	@echo "building program warrior..."
	@go build -o $(WARRIOR_PROG) $(WARRIOR_SRCS)

clean:
	rm -f $(WARRIOR_PROG)