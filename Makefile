.PHONY: migrate-down migrate-up

migrate-down:
	migrate -path internal/infrastructure/migration -database "postgresql://root:secret@localhost:5432/core?sslmode=disable" -verbose down

migrate-up:
	migrate -path internal/infrastructure/migration -database "postgresql://root:secret@localhost:5432/core?sslmode=disable" -verbose up

air:
	air