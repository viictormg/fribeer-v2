package constants

const (
	CompanyIDTest        = "4f76fe5e-527f-11ed-867b-005056a61a3a"
	IDProductTypeProduct = "1ac6e11c-67e5-11ed-867b-005056a61a3a"
	TypePeopleCustomerID = "ffaa49a1-fe4f-11ec-9357-005056a6bad8"
	MessageCreate        = "guardado exitoso"
	MessageErrorCreate   = "error al guardar"
	MessageErrorFind     = "error al consultar"
	MessageNotFound      = "registros no encontrados"
	MessageFound         = "registros  encontrados"
)

var TypeProducts = map[string]string{
	"PROD": "Producto",
	"SERV": "Servicio",
}
