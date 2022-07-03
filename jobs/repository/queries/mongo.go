package queries

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoQuery struct {
}

func (m MongoQuery) GetJobsByFilter(filter map[string]interface{}) interface{} {
	var query bson.D
	title := filter["title"].(string)
	lng := filter["lng"].(float64)
	lat := filter["lat"].(float64)
	radius := filter["radius"].(int64)
	//$minDistance property is set to zero to also include the job at that coordinate and near by radius
	locationQuery := bson.D{{"location", bson.D{{"$near", bson.D{{"$geometry", bson.D{{"type", "Point"}, {"coordinates", bson.A{lng, lat}}}}, {"$minDistance", 0}, {"$maxDistance", radius}}}}}}
	if title != "" {
		query = bson.D{{"$and", bson.A{locationQuery, bson.D{
			{
				Key: "$or", Value: bson.A{
					bson.M{
						"title": bson.M{
							"$regex": primitive.Regex{
								Pattern: "^" + title,
								Options: "i",
							},
						},
					},
					bson.M{
						"title": bson.M{
							"$regex": primitive.Regex{
								Pattern: title + "$",
								Options: "i",
							},
						},
					},
					bson.M{
						"title": bson.M{
							"$regex": primitive.Regex{
								Pattern: title,
								Options: "i",
							},
						},
					},
				},
			},
		}}}}
	} else {
		query = locationQuery
	}
	return query
}
