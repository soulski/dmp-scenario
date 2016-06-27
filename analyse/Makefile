IMAGE_NAME = analyse

build:
	sh -c "'$(CURDIR)/script/build.sh'"
	echo "==> Create docker dmp..."
	docker build -t $(IMAGE_NAME) .

dry-build:
	sh -c "'$(CURDIR)/script/build.sh'"
	echo "==> Create docker $(IMAGE_NAME)..."
	@CID=$$(docker create $(IMAGE_NAME)) && \
		docker cp $(CURDIR)/bin/$(IMAGE_NAME) $$CID:/usr/bin/$(IMAGE_NAME) && \
		docker cp $(CURDIR)/../../dmp/bin/dmp $$CID:/usr/bin/dmp && \
		docker cp $(CURDIR)/supervisord.conf $$CID:/etc/supervisor/conf.d/supervisord.conf && \
		docker stop $$CID && \
		docker commit $$CID $(IMAGE_NAME) && \
		docker rm -vf $$CID

dev: dry-build
	docker run -d --name $(IMAGE_NAME) -p 8080:80 $(IMAGE_NAME)

stop:
	docker rm -vf $(IMAGE_NAME)
