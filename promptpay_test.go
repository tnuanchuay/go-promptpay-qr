package go_promptpay_qr

import (
	"fmt"
	"testing"
)

func TestSanitizePhoneNumber12345678901234ShouldReturnErrInvalidPhoneNumber(t *testing.T){
	_, err := sanitizePhoneNumber("12345678901234", "66")
	if err != ErrInvalidPhoneNumber {
		t.Errorf("expect %v actual %v", ErrInvalidPhoneNumber, err)
	}
}

func TestSanitizePhoneNumber0166800000000_66ShouldReturnErrInvalidPhoneNumber(t *testing.T){
	_, err := sanitizePhoneNumber("0166800000000", "66")
	if err != ErrInvalidPhoneNumber {
		t.Errorf("expect %v actual %v", ErrInvalidPhoneNumber, err)
	}
}

func TestSanitizePhoneNumber110800000000_66ShouldReturnErrInvalidPhoneNumber(t *testing.T){
	_, err := sanitizePhoneNumber("110800000000", "66")
	if err != ErrInvalidPhoneNumber {
		t.Errorf("expect %v actual %v", ErrInvalidPhoneNumber, err)
	}
}

func TestSanitizePhoneNumber11800000000_66ShouldReturnErrInvalidPhoneNumber(t *testing.T){
	_, err := sanitizePhoneNumber("11800000000", "66")
	if err != ErrInvalidPhoneNumber {
		t.Errorf("expect %v actual %v", ErrInvalidPhoneNumber, err)
	}
}

func TestSanitizePhoneNumber1234567890ShouldReturnErrInvalidPhoneNumber(t *testing.T){
	_, err := sanitizePhoneNumber("1234567890", "66")
	if err != ErrInvalidPhoneNumber {
		t.Errorf("expect %v actual %v", ErrInvalidPhoneNumber, err)
	}
}

func TestSanitizePhoneNumber12345ShouldReturnErrInvalidPhoneNumber(t *testing.T) {
	_, err := sanitizePhoneNumber("12345", "66")
	if err != ErrInvalidPhoneNumber {
		t.Errorf("expect %v actual %v", ErrInvalidPhoneNumber, err)
	}
}

func TestSanitizePhoneNumber0066800000000_66ShouldReturn0066800000000(t *testing.T) {
	r, _ := sanitizePhoneNumber("0066800000000", "66")
	e := "0066800000000"
	if r != e {
		t.Errorf("expect %s actual %s", e, r)
	}
}

func TestSanitizePhoneNumber660800000000_66ShouldReturn0066800000000(t *testing.T) {
	r, _ := sanitizePhoneNumber("660800000000", "66")
	e := "0066800000000"
	if r != e {
		t.Errorf("expect %s actual %s", e, r)
	}
}

func TestSanitizePhoneNumber66800000000_66ShouldReturn0066800000000(t *testing.T) {
	r, _ := sanitizePhoneNumber("66800000000", "66")
	e := "0066800000000"
	if r != e {
		t.Errorf("expect %s actual %s", e, r)
	}
}

func TestSanitizePhoneNumber0800000000_66ShouldReturn0066800000000(t *testing.T) {
	r, _ := sanitizePhoneNumber("0800000000", "66")
	e := "0066800000000"
	if r != e {
		t.Errorf("expect %s actual %s", e, r)
	}
}

func TestSanitizeTarget1900000000000ShouldReturn1900000000000(t *testing.T) {
	r := sanitizeTarget("1900000000000")
	e := "1900000000000"
	if r != e {
		t.Errorf("expect %s actual %s", e, r)
	}
}

func TestSanitizeTargetP66800000000shouldReturn66800000000(t *testing.T) {
	r := sanitizeTarget("+66800000000")
	e := "66800000000"
	if r != e {
		t.Errorf("expect %s actual %s", e, r)
	}
}

func AssertField(t *testing.T, a, e string) {
	if e != a {
		t.Errorf("expect %s actual %s", e, a)
	}
}

func TestBuildFormatIndicatorField000201ShouldReturn000201(t *testing.T) {
	AssertField(t,
		buildField(TAG_PAYLOAD_FORMAT_INDICATOR, PAYLOAD_FORMAT_INDICATOR_VALUE),
		"000201")
}

func TestBuildPointOfInitiationFieldMultipleUseShouldReturn010211(t *testing.T) {
	AssertField(t,
		buildField(TAG_POINT_OF_INITIATION_METHOD, MULTIPLE_USE),
		"010211")
}

func TestBuildPointOfInitiationFieldSingleUseShouldReturn010212(t *testing.T) {
	AssertField(t,
		buildField(TAG_POINT_OF_INITIATION_METHOD, SINGLE_USE),
		"010212")
}

func TestBuildFieldMerchantIDWithSubFieldCustomerPresentedQRWithMobileNumber(t *testing.T) {
	AssertField(t,
		buildField(
			TAG_MERCHANT_IDENTIFIER_CREDIT_TRANSFER,
			fmt.Sprintf("%s%s",
				buildField(SUBTAG_MERCHANT_AID, CUSTOMER_PRESENTED_QR),
				buildField(SUBTAG_MOBILE_NUMBER, "0066800000000"))),
		"29370016A00000067701011401130066800000000")
}
