install:
	@chmod +x scripts/*
	cd ./scripts && ./docker-install.sh

setup:
	@chmod +x scripts/*
	cd ./scripts && ./docker-setup.sh
