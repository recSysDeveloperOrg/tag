package model

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

type TagDao struct {
}

type Tag struct {
	ID      string `bson:"_id"`
	Content string `bson:"content"`
}

type TagMovie struct {
	MovieID     string `bson:"movie_id"`
	TagID       string `bson:"tag_id"`
	UpdatedAt   int64  `bson:"updated_at"`
	TaggedTimes int64  `bson:"tagged_times"`
}

type TagUser struct {
	TagID     string `bson:"tag_id"`
	Content   string
	MovieIDs  []string `bson:"movie_ids"`
	UpdatedAt int64    `bson:"updated_at"`
	UserID    string   `bson:"user_id"`
}

var tagDao *TagDao
var tagDaoOnce sync.Once

var ErrDuplicateTag = errors.New("tag exists")

const errDuplicateCode = 11000

func NewTagDao() *TagDao {
	tagDaoOnce.Do(func() {
		tagDao = &TagDao{}
	})

	return tagDao
}

func (*TagDao) InsertTag(ctx context.Context, content string) error {
	_, err := GetClient().Collection(CollectionTag).InsertOne(ctx, struct {
		Content string `bson:"content"`
	}{content})
	if err != nil {
		var writeErr mongo.WriteException
		if errors.As(err, &writeErr) {
			if len(writeErr.WriteErrors) > 0 {
				if writeErr.WriteErrors[0].Code == errDuplicateCode {
					return ErrDuplicateTag
				}
			}
		}

		return err
	}

	return nil
}

func (*TagDao) QueryTagID(ctx context.Context, content string) (*string, error) {
	var tag Tag
	if err := GetClient().Collection(CollectionTag).FindOne(ctx, bson.D{{"content", content}}).
		Decode(&tag); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, err
	}

	return &tag.ID, nil
}

func (*TagDao) IncMovieTag(ctx context.Context, movieID, tagID string) error {
	movieObjectID, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		return err
	}
	tagObjectID, err := primitive.ObjectIDFromHex(tagID)
	if err != nil {
		return err
	}
	if _, err := GetClient().Collection(CollectionTagMovie).UpdateOne(ctx, bson.D{{"movie_id", movieObjectID},
		{"tag_id", tagObjectID}}, bson.D{{"$inc", bson.D{{"tagged_times", 1},
		{"updated_at", time.Now().Unix()}}}},
		options.Update().SetUpsert(true)); err != nil {
		return err
	}

	return nil
}

func (*TagDao) InsertUserTags(ctx context.Context, userID, tagID, movieID string) (bool, error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return false, err
	}
	tagObjectID, err := primitive.ObjectIDFromHex(tagID)
	if err != nil {
		return false, err
	}
	movieObjectID, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		return false, err
	}
	var tagUser TagUser
	if err := GetClient().Collection(CollectionTagUser).FindOne(ctx, bson.D{{"user_id", userObjectID},
		{"tag_id", tagObjectID}}).Decode(&tagUser); err != nil {
		// 用户标签不存在，没关系，等会创建就行了
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return false, err
		}
	}
	for _, existMovieID := range tagUser.MovieIDs {
		if existMovieID == movieID {
			// 用户已经在这个电影上打过这个tag了，不用再加进去了
			return false, nil
		}
	}

	if _, err := GetClient().Collection(CollectionTagUser).UpdateOne(ctx, bson.D{{"user_id", userObjectID},
		{"tag_id", tagObjectID}}, bson.D{{"$push", bson.D{{"movie_ids", movieObjectID}}}},
		options.Update().SetUpsert(true)); err != nil {
		return false, err
	}

	return true, nil
}

func (d *TagDao) QueryMovieTag(ctx context.Context, userID, movieID string) ([]*TagUser, error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	movieObjectID, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		return nil, err
	}
	var tagUser []*TagUser
	c, err := GetClient().Collection(CollectionTagUser).Find(ctx, bson.D{{"user_id", userObjectID},
		{"movie_ids", movieObjectID}}, options.Find().SetSort(bson.D{{"updated_at", -1}}))
	if err != nil {
		return nil, err
	}
	if err := c.All(ctx, &tagUser); err != nil {
		return nil, err
	}
	tagIDs := make([]string, len(tagUser))
	for i, tu := range tagUser {
		tagIDs[i] = tu.TagID
	}
	tags, err := d.QueryTagsByTagID(ctx, tagIDs)
	if err != nil {
		return nil, err
	}
	for i, tag := range tags {
		tagUser[i].Content = tag.Content
	}

	return tagUser, nil
}

func (*TagDao) QueryTagsByTagID(ctx context.Context, tagIDs []string) ([]*Tag, error) {
	tagObjectIDs := make([]primitive.ObjectID, len(tagIDs))
	for i, tagID := range tagIDs {
		tagObjectID, err := primitive.ObjectIDFromHex(tagID)
		if err != nil {
			return nil, err
		}
		tagObjectIDs[i] = tagObjectID
	}

	var tags []*Tag
	c, err := GetClient().Collection(CollectionTag).Find(ctx, bson.D{{"_id", bson.D{{"$in", tagObjectIDs}}}})
	if err != nil {
		return nil, err
	}
	if err := c.All(ctx, &tags); err != nil {
		return nil, err
	}

	return tags, nil
}
