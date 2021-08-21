INSTALL_DIR=/usr/local/bin

build:
	go mod tidy
	go build godaddy.go

install: godaddy
	mv ./godaddy $(INSTALL_DIR)

uninstall:
	rm $(INSTALL_DIR)/godaddy

clean: godaddy
	rm godaddy

.PHONY: build install uninstall clean
