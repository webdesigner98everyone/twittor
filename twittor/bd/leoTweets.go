package bd

import (
	"context"
	// "log"
	"time"

	"github.com/webdesigner98everyone/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoTweets lee los tweets de un perfil */
func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	var resultados []*models.DevuelvoTweets

	/*userid debe ser el mismo que inserte el twwet esta variable esta en insertoTweet.go */
	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	/*Muestra 20 tweets por pagina*/
	opciones.SetLimit(20)
	/*Ordena los twwets por fecha en orden descendente*/
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	/*SetSkip nos ayuda a paginar los twweets por cantidad de twets que son de a 20 por pagina*/
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		// Nos muestra por consola algun error
		// log.Fatal(err.Error())
		return resultados, false
	}

	/* recorre cursor*/
	for cursor.Next(context.TODO()) {

		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}
	return resultados, true
}
