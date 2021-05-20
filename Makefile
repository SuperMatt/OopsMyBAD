.PHONY: install
install: build
	cp bin/oops bin/mybad /usr/local/bin/
	cp systemd/MyBAD.service /etc/systemd/system
	systemctl daemon-reload
	mkdir -p /usr/local/etc/OopsMyBAD/
	cp config_example.yaml /usr/local/etc/OopsMyBAD/config.yaml


.PHONY: build
build: oops mybad

bin:
	mkdir -p bin

.PHONY: mybad
mybad: bin/mybad
bin/mybad: bin
	go build -o bin/mybad .../mybad/

.PHONY: oops
oops: bin/oops
bin/oops: bin
	go build -o bin/oops .../oops/

.PHONY: clean
clean:
	rm -r bin