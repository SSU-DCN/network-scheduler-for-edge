REPO = angel1041502/
NAME = network-scheduler
TAG = 0.0.1

build:
	docker build -t $(REPO)$(NAME):$(TAG) .

push: build
	docker push $(REPO)$(NAME):$(TAG)
