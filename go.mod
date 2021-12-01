module github.com/NpoolPlatform/login-door

go 1.16

require (
	entgo.io/ent v0.9.1
	github.com/NpoolPlatform/application-management v0.0.0-20211130040201-edb8269b3e6f
	github.com/NpoolPlatform/go-service-framework v0.0.0-20211119115808-35513fcc0b81
	github.com/NpoolPlatform/user-management v0.0.0-20211201031609-26a72642b79a
	github.com/NpoolPlatform/verification-door v0.0.0-20211201031400-538c698eeea2
	github.com/casbin/casdoor v1.2.0
	github.com/go-chi/chi/v5 v5.0.5
	github.com/go-redis/redis/v8 v8.11.4
	github.com/go-resty/resty/v2 v2.7.0
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.0
	github.com/thanhpk/randstr v1.0.4
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/net v0.0.0-20211116231205-47ca1ff31462 // indirect
	golang.org/x/oauth2 v0.0.0-20211005180243-6b3c2da341f1 // indirect
	golang.org/x/sys v0.0.0-20211117180635-dee7805ff2e1 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/genproto v0.0.0-20211118181313-81c1377c94b1
	google.golang.org/grpc v1.41.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.41.0
