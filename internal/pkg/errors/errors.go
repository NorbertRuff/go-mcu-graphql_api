package errors

type WrongUsernameOrPasswordError struct{}

func (m *WrongUsernameOrPasswordError) Error() string {
	return "Wrong username or password"
}
