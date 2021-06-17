package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/webdesigner98everyone/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*InsertoTweet graba el Tweet en la BD */
func InsertoTweet(t models.GraboTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	/*registro trae los campos del formulario graboTweet que se ubica en models */
	registro := bson.M{
		"userid":  t.UserID,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}
	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
