install:
	@chmod +x ./scripts/docker-install.sh
	cd ./scripts && ./docker-install.sh

setup:
	@chmod +x ./scripts/docker-setup.sh
	cd ./scripts && ./docker-setup.sh
