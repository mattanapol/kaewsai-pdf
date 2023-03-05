package nosql_repos

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mattanapol/kaewsai-pdf/internal/setting"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SetupMongoDb initializes the NoSQL database instance
func SetupMongoDb(mongoDbSetting setting.MongoDb,
) *mongo.Database {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?retryWrites=false",
		mongoDbSetting.User,
		mongoDbSetting.Password,
		mongoDbSetting.Host,
		mongoDbSetting.Port,
		mongoDbSetting.Name,
	)
	clientOpts := options.Client().ApplyURI(uri).SetMonitor(getMongoCommandMonitor(*log.Default()))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalf("mongo open err: %v", err)
	}

	return client.Database(mongoDbSetting.Name)
}

type MongoConnection struct {
	client mongo.Client
}

func NewMongoDbConnection(mongo *mongo.Database) *MongoConnection {
	return &MongoConnection{client: *mongo.Client()}
}

func (m MongoConnection) Ping() error {
	return m.client.Ping(context.TODO(), nil)
}

func getMongoCommandMonitor(logger log.Logger) *event.CommandMonitor {
	return &event.CommandMonitor{
		Started: func(ctx context.Context, e *event.CommandStartedEvent) {
			logger.Printf("command: %s", e.Command)
		},
		Succeeded: func(ctx context.Context, e *event.CommandSucceededEvent) {
			logger.Printf("reply: %s", e.Reply)
		},
		Failed: func(ctx context.Context, e *event.CommandFailedEvent) {
			logger.Printf("error: %s", e.Failure)
		},
	}
}
