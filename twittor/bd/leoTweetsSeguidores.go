package bd

import (
	"context"
	"time"

	"github.com/webdesigner98everyone/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*LeoTweetsSeguidores lee los tweets de mis seguidores */
func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	// resultados por pagina
	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		// union de tablas en mongo
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		}})

	// CONDICIONES:
	//unwind que todos los documentos nos venga igual
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	// sort nos ordena los resultados en este caso por fecha del mas viejo al mas nuevo a mayor
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweet.fecha": -1}})
	// skip paginacion
	condiciones = append(condiciones, bson.M{"$skip": skip})
	// limit implantamos limites de los resultados en este caso 20 por busqueda
	condiciones = append(condiciones, bson.M{"$limit": 20})

	// Aggregate funcion de mongo nueva, ejecuta internamente en la BD
	cursor, err := col.Aggregate(ctx, condiciones)
	var result []models.DevuelvoTweetsSeguidores
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
