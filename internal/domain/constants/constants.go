package constants

const (
	CompanyIDTest        = "4f76fe5e-527f-11ed-867b-005056a61a3a"
	IDProductTypeProduct = "1ac6e11c-67e5-11ed-867b-005056a61a3a"
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
