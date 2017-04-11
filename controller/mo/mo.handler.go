package mo

import (
	"net/http"
	"time"

	"github.com/Gujarats/send-program/model/global"
	"github.com/Gujarats/send-program/model/mo"
	"github.com/Gujarats/send-program/util"
	"github.com/Gujarats/send-program/util/token"
)

func SendMo(moModel mo.Mo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
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

		elapsed := time.Since(startTime).Seconds()
		global.SetOkResponse(w, "Ok", "Data saved", elapsed, nil)

	})
}

func StatMo(moModel mo.Mo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		date := r.FormValue("date")

		if !util.CheckValue(date) {

			global.SetBadResponse(w, "Failed", "Empty Params")
			return
		}

		result := moModel.GetStats(date)

		global.SetOkResponse(w, "Ok", "Successfully", time.Since(startTime).Seconds(), result)

	})
}
