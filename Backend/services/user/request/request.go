package request

type RequestCreateUser struct {
	UserId   int64 `jsom:"id"`
	UserName int64 `jsom:"username"`
	Email    int64 `jsom:"email"`
	Role     int64 `jsom:"role"`
	Password int64 `jsom:"password"`
}

type RequestUpdate struct {
	UserName int64 `jsom:"username"`
	Email    int64 `jsom:"email"`
	Role     int64 `jsom:"role"`
}

type RequestPasswordChange struct {
	Password int64 `jsom:"password"`
}
