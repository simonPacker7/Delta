module github.com/simonPacker7/Delta/backend/shared/redisclient

go 1.25.4

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
)

require (
	github.com/google/uuid v1.6.0
	github.com/redis/go-redis/v9 v9.17.2
	github.com/simonPacker7/Delta/backend/shared/entities v0.0.0
)

replace github.com/simonPacker7/Delta/backend/shared/entities => ../entities
