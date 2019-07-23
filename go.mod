module github.com/koinotice/oneplus

go 1.12

require (
	github.com/albrow/stringset v2.1.0+incompatible
	github.com/btcsuite/btcd v0.0.0-20190629003639-c26ffa870fd8
	github.com/cevaris/ordered_map v0.0.0-20190319150403-3adeae072e73
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
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
)

// replace github.com/koinotice/oneplus/backend => ../hydro-sdk-backend
