package entity

type PaymentEntity struct {
	ID            string
	TypePayment   string
	PaymentMethod string
	Total         float64
	Observations  string
	Reason        string
	State         string
	Company       string
	createAt      string
}
