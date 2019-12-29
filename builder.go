package go_promptpay_qr

type PromptPayBuilder struct {
	initiationMethod string
	aid              string
	targetType       string
	target           string
	country          string
	currency         string
	amount           float64
}

func NewBuilder() *PromptPayBuilder {
	return &PromptPayBuilder{
		initiationMethod: SINGLE_USE,
		aid:              MERCHANT_PRESENTED_QR,
	}
}

func (p *PromptPayBuilder) Build() (string, error) {
	return Generate(
		p.initiationMethod,
		p.aid,
		p.targetType,
		p.target,
		p.country,
		p.currency,
		p.amount)
}

func (p *PromptPayBuilder) WithAmount(amount float64) *PromptPayBuilder {
	p.amount = amount
	return p
}

func (p *PromptPayBuilder) WithCurrency(code string) *PromptPayBuilder {
	p.currency = code
	return p
}

func (p *PromptPayBuilder) WithCountry(code string) *PromptPayBuilder {
	p.country = code
	return p
}

func (p *PromptPayBuilder) WithBankAccount(number string) *PromptPayBuilder {
	p.target = number
	p.targetType = SUBTAG_BANK_ACCOUNT
	return p
}

func (p *PromptPayBuilder) WithEWallet(id string) (*PromptPayBuilder, error) {
	if len(id) != 15 {
		return nil, ErrInvalidEWallet
	}

	p.target = id
	p.targetType = SUBTAG_EWALLET_ID

	return p, nil
}

func (p *PromptPayBuilder) WithTaxId(id string) (*PromptPayBuilder, error) {
	return p.WithNationalID(id)
}

func (p *PromptPayBuilder) WithNationalID(id string) (*PromptPayBuilder, error) {
	if len(id) != 13 {
		return nil, ErrInvalidNationalId
	}

	p.target = id
	p.targetType = SUBTAG_NATIONAL_TAX_ID

	return p, nil
}

func (p *PromptPayBuilder) WithPhoneNumber(number, countryCode string) (*PromptPayBuilder, error) {
	sanitizedPhoneNumber, err := sanitizePhoneNumber(number, countryCode)
	if err != nil {
		return nil, err
	}

	p.target = sanitizedPhoneNumber
	p.targetType = SUBTAG_MOBILE_NUMBER
	return p, nil
}

func (p *PromptPayBuilder) WithCustomerPresentedQR() *PromptPayBuilder {
	p.aid = CUSTOMER_PRESENTED_QR
	return p
}

func (p *PromptPayBuilder) WithMerchantPresentedQR() *PromptPayBuilder {
	p.aid = MERCHANT_PRESENTED_QR
	return p
}

func (p *PromptPayBuilder) WithMultipleUse() *PromptPayBuilder {
	p.initiationMethod = MULTIPLE_USE
	return p
}

func (p *PromptPayBuilder) WithSingleUse() *PromptPayBuilder {
	p.initiationMethod = SINGLE_USE
	return p
}
