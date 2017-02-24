install:
	@chmod +x utils/scripts/*
	cd ./utils/scripts && ./docker-install.sh

setup:
	@chmod +x utils/scripts/*
	@sudo rm -f /etc/docker/daemon.json
	sudo service docker restart
	cd ./utils/scripts && ./docker-setup.sh

setuptrinity:
	@chmod +x utils/scripts/*
	@sudo rm -f /etc/docker/daemon.json
	echo "{\"dns\": [\"134.226.251.200\", \"134.226.251.100\"]}" | sudo tee -a /etc/docker/daemon.json
	sudo service docker restart
	cd ./utils/scripts && ./docker-setup.sh

clean:
	@sudo rm -f /etc/docker/daemon.json
	sudo service docker restart

