##
# File: Makefile
# Project: gosvelte
# File Created: 05 Jan 2022 22:00:07
# Author: und3fined (me@und3fined.com)
# -----
# Last Modified: 08 Jan 2022 19:22:02
# Modified By: und3fined (me@und3fined.com)
# -----
# Copyright (c) 2022 und3fined.com
##

CLIENT_PATH=client
SERVER_PATH=server

.PHONY: run-client run-server

install-adapter-deps:
	cd $(CLIENT_PATH)/adapter-go && pnpm install

pre-run-client: install-adapter-deps # Install client dependencies
	cd $(CLIENT_PATH) && pnpm install

run-client: pre-run-client
	cd $(CLIENT_PATH) && pnpm dev

build-client: pre-run-client
	cd $(CLIENT_PATH) && pnpm build

run-server:
	cd $(SERVER_PATH) && go run cmd/main/main.go