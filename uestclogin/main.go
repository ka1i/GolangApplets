package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// ! ! ! ! ! !
// use this program ,frist change you account info in blow .

func main() {
	//192.168.9.8  校园网认证ip
	BareURL := "http://192.168.9.8"

	FormData := url.Values{}
	FormData.Add("action", "login")
	FormData.Add("username", "<username>")
	FormData.Add("password", "<password>")
	FormData.Add("ac_id", "1")
	FormData.Add("domain", "@uestc")
	FormData.Add("ajax", "1")

	res, err := http.PostForm(BareURL+"/include/auth_action.php", FormData)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer res.Body.Close()

	bodyContent, err := ioutil.ReadAll(res.Body)
	log.Printf("status code: %d \n", res.StatusCode)
	if strings.Contains(string(bodyContent), "login_ok") {
		log.Printf("login ok")
	} else {
		log.Printf("%s", string(bodyContent))
	}
	return
}
