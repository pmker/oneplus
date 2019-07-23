module github.com/koinotice/oneplus

go 1.12

require (
	github.com/btcsuite/btcd v0.0.0-20190629003639-c26ffa870fd8
	github.com/cevaris/ordered_map v0.0.0-20190319150403-3adeae072e73
	github.com/go-redis/redis v6.15.3+incompatible
	github.com/gorilla/websocket v1.4.0
	github.com/jarcoal/httpmock v1.0.4 // indirect
	github.com/labstack/gommon v0.2.8
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/onrik/ethrpc v0.0.0-20190305112807-6b8e9c0e9a8f
	github.com/petar/GoLLRB v0.0.0-20190514000832-33fb24c13b99
	github.com/satori/go.uuid v1.2.0
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
	github.com/sirupsen/logrus v1.4.1
	github.com/stretchr/testify v1.2.2
	github.com/tidwall/gjson v1.3.2
	github.com/valyala/fasttemplate v1.0.1 // indirect
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
)

// replace github.com/koinotice/oneplus/backend => ../hydro-sdk-backend
