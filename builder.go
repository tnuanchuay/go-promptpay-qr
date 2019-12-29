package go_promptpay_qr

type promptPayBuilder struct {
	initiationMethod string
	aid              string
	targetType       string
	target           string
	country          string
	currency         string
	amount           float64
}

func NewBuilder() *promptPayBuilder {
	return &promptPayBuilder{
		initiationMethod: SINGLE_USE,
		aid:              MERCHANT_PRESENTED_QR,
	}
}

func (p *promptPayBuilder) Build() (string, error) {
	return Generate(
		p.initiationMethod,
		p.aid,
		p.targetType,
		p.target,
		p.country,
		p.currency,
		p.amount)
}

func (p *promptPayBuilder) WithAmount(amount float64) *promptPayBuilder {
	p.amount = amount
	return p
}

func (p *promptPayBuilder) WithCurrency(code string) *promptPayBuilder {
	p.currency = code
	return p
}

func (p *promptPayBuilder) WithCountry(code string) *promptPayBuilder {
	p.country = code
	return p
}

func (p *promptPayBuilder) WithBankAccount(number string) *promptPayBuilder {
	p.target = number
	p.targetType = SUBTAG_BANK_ACCOUNT
	return p
}

func (p *promptPayBuilder) WithEWallet(id string) (*promptPayBuilder, error) {
	if len(id) != 15 {
		return nil, ErrInvalidEWallet
	}

	p.target = id
	p.targetType = SUBTAG_EWALLET_ID

	return p, nil
}

func (p *promptPayBuilder) WithTaxId(id string) (*promptPayBuilder, error) {
	return p.WithNationalID(id)
}

func (p *promptPayBuilder) WithNationalID(id string) (*promptPayBuilder, error) {
	if len(id) != 13 {
		return nil, ErrInvalidNationalId
	}

	p.target = id
	p.targetType = SUBTAG_NATIONAL_TAX_ID

	return p, nil
}

func (p *promptPayBuilder) WithPhoneNumber(number, countryCode string) (*promptPayBuilder, error) {
	sanitizedPhoneNumber, err := sanitizePhoneNumber(number, countryCode)
	if err != nil {
		return nil, err
	}

	p.target = sanitizedPhoneNumber
	p.targetType = SUBTAG_MOBILE_NUMBER
	return p, nil
}

func (p *promptPayBuilder) WithCustomerPresentedQR() *promptPayBuilder {
	p.aid = CUSTOMER_PRESENTED_QR
	return p
}

func (p *promptPayBuilder) WithMerchantPresentedQR() *promptPayBuilder {
	p.aid = MERCHANT_PRESENTED_QR
	return p
}

func (p *promptPayBuilder) WithMultipleUse() *promptPayBuilder {
	p.initiationMethod = MULTIPLE_USE
	return p
}

func (p *promptPayBuilder) WithSingleUse() *promptPayBuilder {
	p.initiationMethod = SINGLE_USE
	return p
}
