build:
	docker-compose build
run:
	docker-compose up -d

migrate:
	docker run --rm --net=chess_net -v \
	$(PWD)/migrations/sql:/flyway/sql \
	flyway/flyway:7.8.1-alpine migrate -url=jdbc:postgresql://postgres:5432/mydb \
	-user=admin -password=qwerty