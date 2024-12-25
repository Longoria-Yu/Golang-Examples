POSTGRESQL_URL=postgresql://user:password@localhost:5432/example?sslmode=disable

# The following `migration` commands refers to:
# https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md
migration-up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

migration-down:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down
