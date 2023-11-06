package migrate

//go:generate go run -mod=mod go.uber.org/mock/mockgen@v0.3.0 -destination=./mocks/migrate.gen.go -package=mocks . Migrator
