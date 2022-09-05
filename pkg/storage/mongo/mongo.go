package mongo

import (
	"GoNews/pkg/storage"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Store Хранилище данных.
type Store struct {
	db *mongo.Client
}

const (
	databaseName   = "data"      // Имя учебной БД
	collectionName = "languages" // Имя коллекции в учебной БД
)

// New Конструктор объекта хранилища.
func New(c string) (*Store, error) {
	mongoOpts := options.Client().ApplyURI(c)
	client, err := mongo.Connect(context.Background(), mongoOpts)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	s := Store{
		db: client,
	}
	return &s, err
}

func (s *Store) Posts() ([]storage.Post, error) {
	collection := s.db.Database(databaseName).Collection(collectionName)
	filter := bson.D{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	var data []storage.Post
	for cur.Next(context.Background()) {
		var l storage.Post
		err := cur.Decode(&l)
		if err != nil {
			return nil, err
		}
		data = append(data, l)
	}
	return data, cur.Err()
}

func (s *Store) AddPost(p storage.Post) error {
	collection := s.db.Database(databaseName).Collection(collectionName)
	//t, _ := strconv.Atoi(time.Now().Format("20060102150405"))
	_, err := collection.InsertOne(context.Background(), p)
	return err
}
func (s *Store) UpdatePost(p storage.Post) error {
	collection := s.db.Database(databaseName).Collection(collectionName)
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"id", p.ID}}
	update := bson.D{{"$set", p}}

	_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	return err
}
func (s *Store) DeletePost(p storage.Post) error {
	collection := s.db.Database(databaseName).Collection(collectionName)
	_, err := collection.DeleteOne(context.TODO(), bson.D{{"id", p.ID}}, options.Delete().SetCollation(&options.Collation{}))
	return err
}
