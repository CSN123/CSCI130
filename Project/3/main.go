package step3

import (


	"net/http"
	"html/template"
	"io"
	"log"
	"github.com/nu7hatch/gouuid"
)

func main() {

	tpl, err := template.ParseFiles("index.html")
	if err != nil {
	log.Fatalln(err)
}

	err = tpl.Execute(res, nil)
	if err != nil {
	log.Fatalln(err)}

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	age := req.FormValue("age")

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
	io.WriteString(res, cookie.Value)



}
