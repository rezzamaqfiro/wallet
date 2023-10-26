package dep

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/rezzamaqfiro/wallet/api"
	"github.com/rezzamaqfiro/wallet/util/buckets"
	"github.com/rezzamaqfiro/wallet/util/buckets/discard"
	"github.com/rezzamaqfiro/wallet/util/buckets/local"
	"github.com/spf13/viper"
)

type DI struct {
	db         *sql.DB
	redis      *redis.Client
	apiHandler http.Handler
}

func InitDI(configFile string) (*DI, error) {
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	di := &DI{}

	return di, nil
}

func (di *DI) GetDatabase() (*sql.DB, error) {
	if di.db == nil {
		db, err := sql.Open("postgres", viper.GetString("database"))
		if err != nil {
			return nil, err
		}

		err = db.Ping()
		if err != nil {
			return nil, err
		}

		di.db = db
	}
	return di.db, nil
}

func (di *DI) GetRedis() (*redis.Client, error) {
	if di.redis == nil {
		rdb := redis.NewClient(&redis.Options{
			Addr:     viper.GetString("redis_host"),
			Password: viper.GetString("redis_pass"), // no password set
			DB:       viper.GetInt("redis_db"),      // use default DB
		})
		di.redis = rdb
	}

	return di.redis, nil
}

func (di *DI) GetBucket() (bucket buckets.Bucket) {
	provider := viper.GetString("bucket.provider")

	switch provider {
	case "local":
		bucket = local.New(
			viper.GetString("bucket.local.path"),
			viper.GetString("server_addr"),
			viper.GetString("bucket.local.url_prefix"),
			nil,
		)

	case "discard":
		bucket = &discard.Bucket{}

	default:
		log.Printf("unknown bucket provider `%v`, using `discard` bucket provider\n", provider)
		bucket = &discard.Bucket{}
	}

	return
}

func (di *DI) GetAPIHandler() (http.Handler, error) {
	if di.apiHandler == nil {
		db, _ := di.GetDatabase()

		bucket := di.GetBucket()
		rdb, _ := di.GetRedis()

		corsHandler := cors.Handler(cors.Options{
			AllowedOrigins: viper.GetStringSlice("cors.allowed_origins"),
			// AllowOriginFunc: func(r *http.Request, origin string) bool { return true },
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			// ExposedHeaders:   []string{"Link"},
			// AllowCredentials: false,
			// MaxAge:           300, // Maximum value not ignored by any of major browsers
		})
		di.apiHandler = api.New(db, rdb, corsHandler).Handler()

		// bucket local server
		if v, ok := bucket.(*local.Bucket); ok {
			v.Handler = di.apiHandler
			di.apiHandler = v
		}
	}

	return di.apiHandler, nil
}

func (di *DI) GetAPIServer() (*http.Server, error) {
	h, err := di.GetAPIHandler()
	if err != nil {
		return nil, err
	}

	srv := http.Server{
		Addr:    viper.GetString("server_addr"),
		Handler: h,
	}
	return &srv, nil
}
