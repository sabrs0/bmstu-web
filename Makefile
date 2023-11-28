build:
	docker compose build
run:
	docker compose up
clean:
	docker container stop bmstu-web-server_main-1 bmstu-web-server1-1 bmstu-web-server2-1  bmstu-web-db_postgres-1
	docker container prune
	docker rmi bmstu-web-server_main:latest
	docker rmi bmstu-web-server1:latest
	docker rmi bmstu-web-server2:latest