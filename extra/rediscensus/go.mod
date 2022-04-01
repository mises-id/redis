module github.com/go-redis/redis/extra/rediscensus/v8

go 1.15

replace github.com/mises-id/redis => ../..

replace github.com/go-redis/redis/extra/rediscmd/v8 => ../rediscmd

require (
	github.com/go-redis/redis/extra/rediscmd/v8 v8.11.5
	github.com/mises-id/redis v8.11.5
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	go.opencensus.io v0.23.0
)
