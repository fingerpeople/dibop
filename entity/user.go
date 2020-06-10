package entity

type ResponseUserDetails struct {
	Code     int         `json:""`
	Meta     *MetaData   `json:"meta"`
	Data     *UserDetail `json:"data"`
	Messages string      `json:"message"`
}

type UserDetail struct {
	AuthData *AuthUsersDetails `json:"authorization"`
	Details  []UserDetailData  `json:"details"`
}

type UserDetailData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Path  string `json:"path"`
}

type AuthUsersDetails struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
