install:
	@chmod +x utils/scripts/*
	cd ./utils/scripts && ./docker-install.sh

setup:
	@chmod +x utils/scripts/*
	@sudo rm -f /etc/docker/daemon.json
	sudo service docker restart

setuptrinity:
	@chmod +x utils/scripts/*
	@sudo rm -f /etc/docker/daemon.json
	echo "{\"dns\": [\"134.226.251.200\", \"134.226.251.100\"]}" | sudo tee -a /etc/docker/daemon.json
	sudo service docker restart

run: 
	$(MAKE) govendor && cd src && go install && src

govendor:
	@chmod +x utils/scripts/vendor.sh && ./utils/scripts/vendor.sh
	@chmod +x utils/scripts/vendor.sh && ./utils/scripts/vendor.sh # Run twice to ensure it's vendored.

test:
	@chmod +x utils/scripts/test.sh && ./utils/scripts/test.sh
