package bookmongorepo

import (
	"context"
	"errors"

	"github.com/yeldisbayev/thrid-group/library/internal/domain"
	"github.com/yeldisbayev/thrid-group/library/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type bookMongoRepo struct {
	collection *mongo.Collection
}

func NewBookMongoRepo(collection *mongo.Collection) repository.BookRepo {
	if collection == nil {
		panic("empty collection argument")
	}

	return bookMongoRepo{
		collection: collection,
	}

}

func (repo bookMongoRepo) Store(
	ctx context.Context,
	book domain.Book,
) (*domain.Book, error) {
	bookSchema := toBookSchema(book)

	result, err := repo.collection.InsertOne(ctx, bookSchema)

	if err != nil {
		return nil, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		book.ID = oid.Hex()
	} else {
		return nil, errors.New("can't get object id")
	}

	return &book, nil

}

func (repo bookMongoRepo) Get(
	ctx context.Context,
	id string,
) (*domain.Book, error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := primitive.M{"_id": objectID}
	bookSchema := Book{}

	result := repo.collection.FindOne(ctx, filter)

	if err := result.Err(); err != nil {
		return nil, err
	}

	if err := result.Decode(&bookSchema); err != nil {
		return nil, err
	}

	book := toBook(bookSchema)

	return &book, nil

}

func (repo bookMongoRepo) Remove(
	ctx context.Context,
	id string,
) error {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := primitive.M{"_id": objectID}

	result, err := repo.collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if result.DeletedCount != 1 {
		return errors.New("can't find document to remove")
	}

	return nil

}
