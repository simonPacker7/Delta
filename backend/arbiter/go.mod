module github.com/simonPacker7/Delta/backend/arbiter

go 1.25.4

require github.com/simonPacker7/Delta/backend/shared/redisclient v0.0.0

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/redis/go-redis/v9 v9.17.2 // indirect
	github.com/simonPacker7/Delta/backend/shared/entities v0.0.0 // indirect
)

replace github.com/simonPacker7/Delta/backend/shared/redisclient => ../shared/redisclient

replace github.com/simonPacker7/Delta/backend/shared/entities => ../shared/entities
