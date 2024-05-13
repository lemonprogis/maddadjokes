.PHONY: all

docker-build:
	docker build -t lemonpro/maddadjokes:beta .

docker-run:
	docker run -it lemonpro/maddadjokes:beta