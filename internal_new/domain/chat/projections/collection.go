package projections

type CollectionProjection struct {
	Chats    map[string]*ChatProjection   `bson:"chats" json:"chats"`
	Messages map[string]*MessageProjction `bson:"messages" json:"messages"`
	Users    map[string]*UserProjection   `bson:"users" json:"users"`

	HasNext bool `bson:"has_next" json:"has_next"`
	HasPrev bool `bson:"has_prev" json:"has_prev"`
}
