package repos

import (
	"crypto/sha256"
	"encoding/hex"
)

// UserIsValid : compare username and password
func UserIsValid(uName, pwd string) bool {

	// Computing sha256 cecksum of the password
	h := sha256.New()
	h.Write([]byte(pwd))
	pwd = hex.EncodeToString(h.Sum(nil))

	// DB simulation (password is tesTLogin)
	_uName, _pwd, _isValid := "medusa", "a3d2ff2ea898335fe81743f963edfa78d9f091faf8cb87bb870630e8c7420a8d", false

	if uName == _uName && pwd == _pwd {
		_isValid = true
	} else {
		_isValid = false
	}

	return _isValid
}
