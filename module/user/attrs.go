package user

type UserModelAttrFunc func(u *Model)

type UserModelAttrFuncs []UserModelAttrFunc

func WithUserId(id int64) UserModelAttrFunc {
	return func(u *Model) {
		u.Id = id
	}
}

func WithUsername(name string) UserModelAttrFunc {
	return func(u *Model) {
		u.Username = name
	}
}

func WithPassword(pwd string) UserModelAttrFunc {
	return func(u *Model) {
		u.Password = pwd
	}
}

func WithEmail(email string) UserModelAttrFunc {
	return func(u *Model) {
		u.Email = email
	}
}

func (this UserModelAttrFuncs) apply(u *Model) {
	for _, f := range this {
		f(u)
	}
}
