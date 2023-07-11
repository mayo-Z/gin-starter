package public

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func RandomUid() uint {
	rand.NewSource(time.Now().Unix())
	return 100000 + uint(rand.Intn(900000))
}

func ValidPassword(afterPw, beforePw string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(afterPw), []byte(beforePw)); err != nil {
		return false
	} else {
		return true
	}
}

func SetHashedPassword(pw string) (string, error) {
	hashPw, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(hashPw), err
}

func TimeToStr(t any) string {

	switch t.(type) {
	case gorm.DeletedAt:
		if t.(gorm.DeletedAt).Valid {
			return t.(gorm.DeletedAt).Time.Format("2006-01-02 15:04:05")
		} else {
			return ""
		}
	case time.Time:
		return t.(time.Time).Format("2006-01-02 15:04:05")
	default:
		return ""
	}
}
