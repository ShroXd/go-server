package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Books struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	BookId string `bson:"bookId,omitempty"`
	BookName string `bson:"bookName,omitempty"`
	BookSimpleDescribe string `bson:"bookSimpleDes,omitempty"`
	BookCategory string `bson:"bookCategory,omitempty"`
	BookWordCount string `bson:"bookWordCount,omitempty"`
	BookImageURL string `bson:"imageUrl,omitempty"`
}
