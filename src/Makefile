
servers = world login front chat
.PHONY: clear
clear:
	for server in $(servers); do \
  		rm -rf ../bin/$$server; \
		echo $$server; \
	done
	rm -rf ../bin/logs

.PHONY: build
build:
	for server in $(servers); do \
  		go build -o ../bin/$$server ./cmd/$$server/main.go;\
		echo $$server; \
	done

.PHONY: logs
logs:
	for server in $(servers); do \
  		mkdir -p ../bin/logs/$$server ;\
		echo $$server; \
	done

.PHONY: startup
startup:logs
	./build/tartup.sh

.PHONY: shutdown
shutdown:
	./build/shutdown.sh
