// Управление сессиями
package session

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"strconv"
	"strings"
)

var secret []byte

// Init загружает секрет из .env в переменную secret
func Init() error {
	s := os.Getenv("SECRET")
	if s == "" {
		return errors.New("SECRET not set in .env")
	}

	secret = []byte(s)

	return nil
}

// Sign возвращает строку состоящую из id, "."
// и случайного набора символов с помощью HMAC.
func Sign(id int) string {
	payload := strconv.Itoa(id)
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(payload))
	signature := h.Sum(nil)

	return payload + "." + hex.EncodeToString(signature)
}

// Verify возвращает id пользователя и boolen: true,
// если куки не подделан и false, если куки подделан
// и не является реальным.
func Verify(value string) (int, bool) {
	str := strings.SplitN(value, ".", 2)
	if len(str) < 2 {
		return 0, false
	}

	payload := str[0]
	sign, err := hex.DecodeString(str[1])
	if err != nil {
		return 0, false
	}

	h := hmac.New(sha256.New, secret)
	h.Write([]byte(payload))
	signature := h.Sum(nil)
	// используем hmac.Equal вместо ==
	// для одинакового времени проверки
	// и безопасности данных.
	if !hmac.Equal(sign, signature) {
		return 0, false
	}

	id, err := strconv.Atoi(payload)
	if err != nil {
		return 0, false
	}

	return id, true
}
