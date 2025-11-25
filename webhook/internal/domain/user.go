package domain

import "time"

type User struct {
	Id       int64
	Email    string
	Password string
	// UTC 0 的时区
	// 这里用的是time.time  而dao 里面的user 用的int64 因为这里负责转化时间戳 所以用time
	//dao 用int 是因为保证时间一致性
	Ctime time.Time
}
