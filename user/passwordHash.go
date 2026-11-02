// Пакет работы с пользователями
package user

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash возвращает слайс байтов хэша пароля, сделанного
// через bcrypt и ошибку если что то пошло не так.
func Hash(password string) ([]byte, error) {
	var err error
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return passwordHash, err
}

// Compare возвращает ошибку если данный хэш и
// пароль не совпадают и ничего если они совпали.
func Compare(passHash string, password string) error {
	var err error
	err = bcrypt.CompareHashAndPassword([]byte(passHash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
