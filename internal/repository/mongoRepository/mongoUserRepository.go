package mongoRepository

import (
	"context"
	"log"

	"github.com/brandoyts/go-clean/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const collection = "users"

type UserMongoRepository struct {
	collection *mongo.Collection
}

func NewUserMongoRepository(database *mongo.Database) domain.UserRepository {
	return &UserMongoRepository{
		collection: database.Collection(collection),
	}
}
func (u *UserMongoRepository) All(ctx context.Context) ([]domain.User, error) {
	cursor, err := u.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var results []domain.User

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
func (u *UserMongoRepository) Find(ctx context.Context, user domain.User) ([]domain.User, error) {
	return nil, nil
}
func (u *UserMongoRepository) FindById(ctx context.Context, id string) (*domain.User, error) {
	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var result domain.User

	filter := bson.D{{"_id", objectId.Hex()}}
	log.Println(filter)

	err = u.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
func (u *UserMongoRepository) Create(ctx context.Context, user domain.User) (string, error) {
	user.ID = bson.NewObjectID().Hex()
	result, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(string), nil
}
func (u *UserMongoRepository) Update(ctx context.Context, user domain.User) error {
	return nil
}
func (u *UserMongoRepository) Delete(ctx context.Context, id string) error {
	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", objectId.Hex()}}

	result := u.collection.FindOneAndDelete(ctx, filter)

	if result.Err() != nil {
		return result.Err()
	}

	return nil
}
