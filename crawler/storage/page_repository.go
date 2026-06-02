package storage

import (
	"context"
	"time"

	"github.com/aditya-sutar-45/search--/crawler/parser"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type PageRepository struct {
	collection *mongo.Collection
}

func NewPageRepository(conn *MongoConnection) (*PageRepository, error) {
	repo := &PageRepository{
		collection: conn.DB.Collection("raw_pages"),
	}

	err := repo.createIndexes()
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *PageRepository) Save(p *parser.Page) error {
	page := pageToPageDocument(p)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, page)
	return err
}

func (r *PageRepository) SaveOrUpdate(p *parser.Page) error {
	page := pageToPageDocument(p)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"url_hash": page.URLHash,
	}

	update := bson.M{
		"$set": page,
	}

	opts := options.UpdateOne().SetUpsert(true)

	_, err := r.collection.UpdateOne(
		ctx,
		filter,
		update,
		opts,
	)

	return err
}

func (r *PageRepository) createIndexes() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys: bson.D{
				{Key: "url_hash", Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	)

	return err
}
