package projections

type ChatProjection struct {
	ID   string `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name"`
	Type string `bson:"type" json:"type"`
}
