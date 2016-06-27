IMAGE_NAME = scenario

build:
	echo "==> Create docker scenario..."
	docker build -t $(IMAGE_NAME) .

	echo "==> Build Gateway..."
	$(MAKE) -C gateway build
	echo "==> Build User..."
	$(MAKE) -C user build
	echo "==> Build Store..."
	$(MAKE) -C store build
	echo "==> Build Authen..."
	$(MAKE) -C authen build
	echo "==> Build Analyse..."
	$(MAKE) -C analyse build
	echo "==> Build Search..."
	$(MAKE) -C search build
	echo "==> Build Logging..."
	$(MAKE) -C logging build

dry-build:
	echo "==> Build Gateway..."
	$(MAKE) -C gateway dry-build
	echo "==> Build User..."
	$(MAKE) -C user dry-build
	echo "==> Build Store..."
	$(MAKE) -C store dry-build
	echo "==> Build Authen..."
	$(MAKE) -C authen dry-build
	echo "==> Build Analyse..."
	$(MAKE) -C analyse dry-build
	echo "==> Build Search..."
	$(MAKE) -C search dry-build
	echo "==> Build Logging..."
	$(MAKE) -C logging dry-build

start:
	docker-compose up

stop:
	docker-compose down
