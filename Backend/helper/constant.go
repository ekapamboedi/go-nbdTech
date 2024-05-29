package helper

type RoleBody struct {
	Id   int64
	Role string
}

var Role = []RoleBody{
	{
		Id:   1,
		Role: "SuperAdmin",
	},
	{
		Id:   2,
		Role: "Admin",
	},
	{
		Id:   3,
		Role: "User",
	},
}
