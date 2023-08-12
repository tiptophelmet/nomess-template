package response

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tiptophelmet/nomess/intl"

	"github.com/BurntSushi/toml"
)

type Response struct {
	Resp string
	Txt  string
}

type ErrorResponses struct {
	ERR map[string]Response
}

type OkResponses struct {
	OK map[string]Response
}

var er *ErrorResponses
var or *OkResponses

func parseResponses(filename string, responses any) {
	tomlData, err := os.ReadFile(fmt.Sprintf("../locales/%s/%s.toml", intl.GetLocale(), filename))
	if err != nil {
		panic(err)
	}

	if _, err := toml.Decode(string(tomlData), &responses); err != nil {
		panic(err)
	}
}

func initResponses() {
	// 1. ERR responses
	parseResponses("err", er)

	// 2. OK responses
	parseResponses("ok", or)
}

func Init(w *http.ResponseWriter) {
	initResponder(w)
	initResponses()
}

func Error(name string, statusCode int) {
	Respond(er.ERR[name], statusCode)
}

func Ok(name string, statusCode int) {
	Respond(or.OK[name], statusCode)
}
