package contracts

type Users struct {
	Id       int64  `bson:"id"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}
