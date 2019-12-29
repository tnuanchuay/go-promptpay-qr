package go_promptpay_qr

import "testing"

func TestBuilder(t *testing.T){
	builder, _ := NewBuilder().
		WithMultipleUse().
		WithMerchantPresentedQR().
		WithCountry(THAILAND).
		WithCurrency(THAI_BAHT).
		WithPhoneNumber("0000000000", "66")

	r, _ := builder.Build()

	e := "00020101021129370016A000000677010111011300660000000005802TH530376463048956"

	if r != e {
		t.Errorf("expect %s actual %s", e, r)
	}
}
