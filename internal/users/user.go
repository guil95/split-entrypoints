package users

type User struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	Age  int64  `json:"age" bson:"age"`
}
