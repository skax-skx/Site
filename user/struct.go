package user

// Структура пользователя из бд.
type User struct {
	Id            int
	Username      string
	Password_hash string
}
