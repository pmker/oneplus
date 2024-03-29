module github.com/pmker/oneplus

go 1.12

require (
	github.com/albrow/stringset v2.1.0+incompatible
	github.com/btcsuite/btcd v0.0.0-20190629003639-c26ffa870fd8
	github.com/cevaris/ordered_map v0.0.0-20190319150403-3adeae072e73
	github.com/dgraph-io/badger v1.6.0-rc1
	github.com/ethereum/go-ethereum v1.9.0
	github.com/go-redis/redis v6.15.3+incompatible
	github.com/google/uuid v1.1.1
	github.com/gorilla/websocket v1.4.0
	github.com/koinotice/vedex v0.0.0-20190723102702-da200cc08436
	github.com/labstack/gommon v0.2.9
	github.com/onrik/ethrpc v1.0.0
	github.com/petar/GoLLRB v0.0.0-20190514000832-33fb24c13b99
	github.com/satori/go.uuid v1.2.0
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.3.0
	github.com/syndtr/goleveldb v1.0.0
	github.com/tidwall/gjson v1.3.2
	go.etcd.io/bbolt v1.3.3
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
)

replace (
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2
	github.com/golang/lint => github.com/golang/lint v0.0.0-20190227174305-8f45f776aaf1
	github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1

	github.com/minio/minio-go => github.com/minio/minio-go v6.0.14+incompatible

	github.com/nats-io/go-nats => github.com/nats-io/nats.go v1.8.1
	github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20181115231424-8e868ca12c0f
	golang.org/x/build => github.com/golang/build v0.0.0-20190403045414-85a73d7451e7

	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190325154230-a5d413f7728c
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190402192236-7fd597ecf556
	golang.org/x/image => github.com/golang/image v0.0.0-20190321063152-3fc05d484e9f
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190313153728-d0100b6bd8b3
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190327163128-167ebed0ec6d
	golang.org/x/net => github.com/golang/net v0.0.0-20190328230028-74de082e2cca
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190402181905-9f3314589c9a
	golang.org/x/perf => github.com/golang/perf v0.0.0-20190312170614-0655857e383f
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190227155943-e225da77a7e6
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190402142545-baf5eb976a8c
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190402200628-202502a5a924
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.3.0
	google.golang.org/appengine => github.com/golang/appengine v1.5.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190401181712-f467c93bbac2
	google.golang.org/grpc => github.com/grpc/grpc-go v1.19.1
)
