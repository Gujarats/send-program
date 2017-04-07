package mo

type MoInterface interface {
	InsertData(msisdn, operatorid, shortcodeid, text, token string) error
}
