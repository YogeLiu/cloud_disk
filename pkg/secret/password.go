package secret

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/YogeLiu/CloudDisk/pkg/util"
)

func SetPassword(plaintext string) (string, error) {
	salt := util.RandStringRunes(16)
	hash := sha1.New()
	_, err := hash.Write([]byte(plaintext + salt))
	bs := hex.EncodeToString(hash.Sum(nil))
	if err != nil {
		return "", err
	}
	password := salt + ":" + string(bs)
	return password, nil
}

func CheckPassword(password, enPassword string) (bool, error) {
	passwordStore := strings.Split(enPassword, ":")
	if len(passwordStore) != 2 {
		return false, errors.New("unknown password type")
	}
	hash := sha1.New()
	_, err := hash.Write([]byte(password + passwordStore[0]))
	bs := hex.EncodeToString(hash.Sum(nil))
	if err != nil {
		return false, err
	}
	return bs == passwordStore[1], nil
}
