module github.com/scalarorg/xchains-api

go 1.23.3

require (
	github.com/btcsuite/btcd v0.24.2
	github.com/btcsuite/btcd/btcec/v2 v2.3.2
	github.com/btcsuite/btcd/btcutil v1.1.6
	github.com/btcsuite/btcd/chaincfg/chainhash v1.1.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/scalarorg/bitcoin-vault/go-utils v0.0.0-20250116044226-07815838b26d
	github.com/scalarorg/data-models v0.0.0-20250123023305-d1c76e395e77
	github.com/scalarorg/xchains-indexer v0.0.0-20241205021823-db0507b3a4ad
	github.com/shopspring/decimal v1.4.0
	github.com/spf13/viper v1.18.2
	github.com/stretchr/testify v1.10.0
	github.com/swaggo/swag v1.16.3
	github.com/unrolled/secure v1.14.0
	gorm.io/driver/postgres v1.5.9
	gorm.io/gorm v1.25.12
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/decred/dcrd/crypto/blake256 v1.0.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/spec v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/snappy v0.0.5-0.20220116011046-fa5810519dcb // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/klauspost/compress v1.17.7 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.52.2 // indirect
	github.com/prometheus/procfs v0.13.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/swaggo/files v0.0.0-20220610200504-28940afbdbfe // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/tools v0.21.1-0.20240508182429-e35e4ccd0d2d // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

require (
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-chi/chi v1.5.5
	github.com/google/uuid v1.6.0
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/joho/godotenv v1.5.1
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.0 // indirect
	github.com/prometheus/client_golang v1.19.0
	github.com/rs/cors v1.10.1
	github.com/rs/zerolog v1.32.0
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/cobra v1.8.0
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/swaggo/http-swagger v1.3.4
	go.mongodb.org/mongo-driver v1.14.0
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20240404231335-c0f41cb1a7a0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

// replace github.com/scalarorg/xchains-indexer => ../xchains-indexer
