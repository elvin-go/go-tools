package mongoutil

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMongoTx(ctx context.Context, client *mongo.Client) (tx.Tx, error) {}
