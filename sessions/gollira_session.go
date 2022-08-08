package sessions

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"io"
	"time"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte(makeID()))

type Cuser struct {
	Name       string
	LoginTime  time.Time
	LogoutTime time.Time
}

func init() {
	gob.Register(&Cuser{})
}

func makeID() string {
	b := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
