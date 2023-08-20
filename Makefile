.PHONY: postgres adminer migrate

postgres:
	#docker run --rm -ti --network host -e POSTGRES_PASSWORD=secret postgres
	docker run --rm -ti -p 5432:5432 -e POSTGRES_PASSWORD=secret postgres

adminer:
	docker run --rm -ti -p 8080:8080 adminer
	# connect via server=192.168.0.51

migrate: 
	migrate -source file://migrations \
		-database postgres://postgres:secret@localhost/postgres?sslmode=disable up
	
migrate-down: 
	migrate -source file://migrations \
		-database postgres://postgres:secret@localhost/postgres?sslmode=disable down
