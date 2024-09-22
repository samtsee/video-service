postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root video_service

dropdb:
	docker exec -it postgres dropdb video_service

.PHONY: postgres createdb dropdb
