package mo

import (
	"log"
	"net/http"
	"os"

	"github.com/Gujarats/send-program/model/global"
	"github.com/Gujarats/send-program/model/mo"
	"github.com/Gujarats/send-program/util"
	"github.com/Gujarats/send-program/util/token"
)

// create logger to print error in the console
var logger *log.Logger

func init() {
	logger = log.New(os.Stderr,
		"Controller Driver :: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func SendMo(moModel mo.Mo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msisdn := r.FormValue("msisdn")
		operatorId := r.FormValue("operatorid")
		shortCodeId := r.FormValue("shortcodeid")
		text := r.FormValue("text")

		if !util.CheckValue(msisdn, operatorId, shortCodeId, text) {

			global.SetBadResponse(w, "Failed", "Empty Params")
			return
		}

		moModel.Msisdn = msisdn
		moModel.OperatorId = operatorId
		moModel.ShortCodeID = shortCodeId
		moModel.Text = text

		// save the data using go routine
		go func(moModel mo.Mo) {
			token, _ := token.GenerateToken(r)
			moModel.InsertData(token)
		}(moModel)

		global.SetOkResponse(w, "Ok", "Data saved", nil)

	})
}
