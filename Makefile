migrate_diff:
	atlas migrate diff --env gorm

migrate_apply:
	atlas migrate apply --url postgres://postgres:postgres@localhost:45432/test?sslmode=disable --dir file://migrations

migrate_status:
	atlas migrate status --url postgres://postgres:postgres@localhost:45432/test?sslmode=disable --dir file://migrations

migrate_hash:
	atlas migrate hash --dir file://migrations

run_test:
	go test -v ./...

clean_test:
	go clean -testcache
