package utility

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func HashStr(trg, alg string) string {
	hashed := ""
	b := []byte(trg)
	switch alg {
	case "md5":
		md5 := md5.Sum(b)
		hashed = hex.EncodeToString(md5[:])
	case "sha1":
		sha1 := sha1.Sum(b)
		hashed = hex.EncodeToString(sha1[:])
	case "sha512":
		sha512 := sha512.Sum512(b)
		hashed = hex.EncodeToString(sha512[:])
	default:
		sha256 := sha256.Sum256(b)
		hashed = hex.EncodeToString(sha256[:])
	}
	return hashed
}
