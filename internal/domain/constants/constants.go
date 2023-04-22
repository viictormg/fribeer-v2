package constants

const (
	CompanyIDTest        = "4f76fe5e-527f-11ed-867b-005056a61a3a"
	IDProductTypeProduct = "1ac6e11c-67e5-11ed-867b-005056a61a3a"
	TypePeopleCustomerID = "ffaa49a1-fe4f-11ec-9357-005056a6bad8"
	MessageCreate        = "guardado exitoso"
	MessageErrorCreate   = "error al guardar"
	MessageErrorFind     = "error al consultar"
	MessageNotFound      = "registros no encontrados"
	MessageFound         = "registros encontrados"
	MessageFoundSingular = "registro encontrado"
	CampusIDTest         = "94f9df9f-0fbd-11ed-9357-005056a6bad8"
	DateFormatInvalid    = "the date format should be 'yyyy-mm-dd hh:mm:ss'"
	TypeProductService   = "Servicio"
)

var TypeProducts = map[string]string{
	"PROD": "Producto",
	"SERV": "Servicio",
}

var StateSaleConstant = map[string]string{
	"PENDIENTE_DE_PAGO": "6a21bb1c-2b48-11ed-860e-005056a61a3a",
	"BORRADOR":          "9d1b5176-2b48-11ed-860e-005056a61a3a",
	"PAGADA":            "fa78d8ea-2b47-11ed-860e-005056a61a3a",
}
