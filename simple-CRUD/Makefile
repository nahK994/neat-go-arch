db_name=user_db
user=user
password=password
host=localhost
port=5432
db_url=postgres://$(user):$(password)@$(host):$(port)/$(db_name)?sslmode=disable


run:
	cd cmd && \
	go run main.go
kill:
	sudo kill -9 `sudo lsof -t -i:8000`


db-up:
	docker compose -f docker-compose.yml up -d
db-down:
	docker compose -f docker-compose.yml down


migrate-up:
	cd pkg/repository/migrations && \
	goose postgres ${db_url} up
migrate-down:
	cd pkg/repository/migrations && \
	goose postgres ${db_url} down
