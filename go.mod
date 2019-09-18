module github.com/loomnetwork/blockatlas

go 1.12

require (
	github.com/btcsuite/btcutil v0.0.0-20190425235716-9e5f4b9a998d
	github.com/deckarep/golang-set v1.7.1
	github.com/gin-gonic/gin v1.4.0
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/mitchellh/mapstructure v1.1.2
	github.com/mr-tron/base58 v1.1.2
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cast v1.3.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.3.0
	github.com/thoas/go-funk v0.4.0
	github.com/trustwallet/blockatlas v1.0.9
	github.com/ybbus/jsonrpc v2.1.2+incompatible
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/trustwallet/blockatlas v1.0.9 => ../blockatlas
