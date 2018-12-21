package ocm

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"log"
	"os"
	"strings"
)

type Config struct {
	username string
	password string
	domain   string
}

func (c Config) Client() (string, bool) {

	type Payload struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}

	basicAuth := Payload{User: c.username, Password: c.password}
	urlAuth := fmt.Sprintf("http://%s/authenticate/", c.domain)

	resp, body, errs := gorequest.New().Post(urlAuth).Set("Content-Type", "application/oracle-compute-v3+json").Send(basicAuth).End()
	if resp.StatusCode != 204 {
		return "meh", false
		log.Println("[ERROR] Problema con:", resp, body, errs)
		os.Exit(1)
	}

	parsedCookie := fmt.Sprint(between(fmt.Sprint(resp), "nimbula", " Max-Age"))
	cookie := fmt.Sprintf("nimbula%s Max-Age=1800", parsedCookie)
	os.Setenv("COMPUTE_COOKIE", cookie)
	return c.domain, true
}

func between(value string, a string, b string) string {
	posFirst := strings.Index(value, a)
	if posFirst == 1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == 1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}
