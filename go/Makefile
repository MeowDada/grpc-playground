.PHONY: all clean

BINARY_DIR=bin

WARRIOR_PROTO_SRCDIR=proto/warrior
WARRIOR_PROTO_DSTDIR=proto/warrior
WARRIOR_PROG=$(BINARY_DIR)/warrior
WARRIOR_SRCS=warrior.go

ECHO_PROTO_SRCDIR=proto/echo
ECHO_PROTO_DSTDIR=proto/echo
ECHO_CLIENT_PROG=$(BINARY_DIR)/echo-client
ECHO_CLIENT_SRCS=echo-client.go
ECHO_SERVER_PROG=$(BINARY_DIR)/echo-server
ECHO_SERVER_SRCS=echo-server.go

all: build_warrior build_echo

build_warrior:
	@echo "generating warrior protobuf GO codes..."
	@protoc -I=$(WARRIOR_PROTO_SRCDIR) --go_out=$(WARRIOR_PROTO_DSTDIR) $(WARRIOR_PROTO_SRCDIR)/*.proto
	@echo "building program warrior..."
	@go build -o $(WARRIOR_PROG) $(WARRIOR_SRCS)

build_echo:
	@echo "generating echo protobuf GO codes..."
	@protoc -I=$(ECHO_PROTO_SRCDIR) --go_out=$(ECHO_PROTO_DSTDIR) $(ECHO_PROTO_SRCDIR)/*.proto
	@echo "building program echo client and server..."
	@go build -o $(ECHO_CLIENT_PROG) $(ECHO_CLIENT_SRCS)
	@go build -o $(ECHO_SERVER_PROG) $(ECHO_SERVER_SRCS)

clean:
	rm -rf $(BINARY_DIR)/*