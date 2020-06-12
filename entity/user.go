package entity

// ResponseUserDetails ...
type ResponseUserDetails struct {
	Code     int         `json:""`
	Meta     *MetaData   `json:"meta"`
	Data     *UserDetail `json:"data"`
	Messages string      `json:"message"`
}

// ResponseUserList ...
type ResponseUserList struct {
	Code     int            `json:""`
	Meta     *MetaData      `json:"meta"`
	Data     []UserListData `json:"data"`
	Messages string         `json:"message"`
}

// UserDetail ...
type UserDetail struct {
	AuthData *AuthUsersDetails `json:"authorization"`
	Details  []UserDetailData  `json:"details"`
}

// UserListData ...
type UserListData struct {
	User  string `json:"user"`
	Value string `json:"role"`
	Key   string `json:"key"`
}

// UserDetailData ...
type UserDetailData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Path  string `json:"path"`
}

// AuthUsersDetails ...
type AuthUsersDetails struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// UserRequest ...
type UserRequest struct {
	User     string `json:"username"`
	Password string `json:"password"`
}
