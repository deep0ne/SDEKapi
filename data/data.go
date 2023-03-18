package data

const (
	ApiURL      = "https://api.edu.cdek.ru/v2/oauth/token"
	CalcURL     = "https://api.edu.cdek.ru/v2/calculator/tarifflist"
	AddressFrom = "Россия, г. Москва, Cлавянский бульвар д.1"
	AddressTo   = "Россия, Воронежская обл., г. Воронеж, ул. Ленина д.43"
)

type Size struct {
	Weight int
	Height int
	Length int
	Width  int
}

type PriceSending struct {
	TariffCode   int     `json:"tariff_code"`
	TariffName   string  `json:"tariff_name"`
	TariffDesc   string  `json:"tariff_description"`
	DeliveryMode int     `json:"delivery_mode"`
	DeliverySum  float64 `json:"delivery_sum"`
	PeriodMin    int     `json:"period_min"`
	PeriodMax    int     `json:"period_max"`
}

type PriceSendings struct {
	TariffCodes []PriceSending `json:"tariff_codes"`
}

func NewPackage(weight, height, length, width int) Size {
	return Size{
		Weight: weight,
		Height: height,
		Length: length,
		Width:  width,
	}
}
