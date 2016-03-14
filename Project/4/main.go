ppackage step4

import (
	"net/http"
	"encoding/json"
	"encoding/base64"
	"log"
)

type User struct {
	Name string
	Age string
}

func user(req *http.Request, user *User) string {
	user.Name = req.FormValue("name")
	user.Age = req.FormValue("age")

	bs, _ := json.Marshal(user)

	b64 := base64.URLEncoding.EncodeToString(bs)


	return b64
}

func index(res http.ResponseWriter, req *http.Request)
	cookie, err := req.Cookie("session-info")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id.String() + "|" + name + "|" + age,
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}

	user := new(User)
	b64 := set_user(req, user)

	if req.FormValue("user_name") != "" {
		cookie_id := strings.Split(cookie.Value, "|")
		cookie.Value = cookie_id[0] + "|" + b64
}
	http.SetCookie(res, cookie)


}
func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
