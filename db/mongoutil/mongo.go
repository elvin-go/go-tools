package mongoutil

import (
	"context"
	"go-tools/db/tx"
	"go-tools/errs"
	"go-tools/mw/specialerror"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func init() {
	if err := specialerror.AddReplace(mongo.ErrNoDocuments, errs.ErrRecordNotFound); err != nil {
		panic(err)
	}
}

type Config struct {
	Uri         string
	Address     []string
	Database    string
	Username    string
	Password    string
	MaxPoolSize int
	MaxRetry    int
}

type Client struct {
	tx tx.Tx
	db *mongo.Database
}

func (c *Client) GetDB() *mongo.Database {
	return c.db
}

func (c *Client) GetTx() tx.Tx {
	return c.tx
}

func NewMongoDB(ctx context.Context, cfg *Config) (*Client, error) {

	credential := options.Credential{
		Username: cfg.Username,
		Password: cfg.Password,
	}

	opts := options.Client().ApplyURI(cfg.Uri).SetAuth(credential).SetMaxPoolSize(uint64(cfg.MaxPoolSize))
	var (
		client *mongo.Client
		err    error
	)

	for i := 0; i < cfg.MaxRetry; i++ {
		client, err = connectMongo(ctx, opts)
		if err != nil && shouldRetry(ctx, err) {
			time.Sleep(time.Second / 2)
			continue
		}
		break
	}

	if err != nil {
		return nil, errs.WrapMsg(err, "failed to connect to MongoDB", "URI", cfg.Uri)
	}
	mtx, err := NewMongoTx(ctx, client)
	if err != nil {
		return nil, err
	}
	return &Client{
		tx: mtx,
		db: client.Database(cfg.Database),
	}, nil
}

func connectMongo(ctx context.Context, opts *options.ClientOptions) (*mongo.Client, error) {
	cli, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	if err := cli.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return cli, nil
}
