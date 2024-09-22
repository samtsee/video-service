postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root video_service

dropdb:
	docker exec -it postgres dropdb video_service

migrateup:
	migrate -path migrations -database "postgres://root:secret@localhost:5432/video_service?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgres://root:secret@localhost:5432/video_service?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown

