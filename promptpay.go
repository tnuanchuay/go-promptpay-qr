package go_promptpay_qr

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	TAG_PAYLOAD_FORMAT_INDICATOR               = "00"
	TAG_POINT_OF_INITIATION_METHOD             = "01"
	TAG_MERCHANT_IDENTIFIER_CREDIT_TRANSFER    = "29"
	TAG_MERCHANT_IDENTIFIER_BILL_PAYMENT       = "30"
	TAG_MERCHANT_IDENTIFIER_PAYMENT_INNOVATION = "31"
	TAG_TRANSACTION_AMOUNT                     = "54"
	TAG_COUNTRY                                = "58"
	TAG_CRC                                    = "63"

	SUBTAG_MERCHANT_AID    = "00"
	SUBTAG_MOBILE_NUMBER   = "01"
	SUBTAG_NATIONAL_TAX_ID = "02"
	SUBTAG_EWALLET_ID      = "03"
	SUBTAG_BANK_ACCOUNT    = "04"

	PAYLOAD_FORMAT_INDICATOR_VALUE = "01"
	SINGLE_USE                     = "12"
	MULTIPLE_USE                   = "11"
	MERCHANT_PRESENTED_QR          = "A000000677010111"
	CUSTOMER_PRESENTED_QR          = "A000000677010114"
)

var (
	ErrInvalidInitiationMethod                 = errors.New("initiation method is invalid")
	ErrInvalidMerchantIdentifierCreditTransfer = errors.New("merchant identifier credit transfer(aid) is invalid")
	ErrInvalidPhoneNumber                      = errors.New("phone number is invalid")
)

//func generate(
//	initiationMethod,
//	aid,
//	targetType,
//	target string,
//	amount float64) (string, error) {
//
//	if isNotInitiationMethod(initiationMethod) {
//		return "", ErrInvalidInitiationMethod
//	}
//
//	if isNotAID(aid) {
//		return "", ErrInvalidMerchantIdentifierCreditTransfer
//	}
//
//	data := []string{
//		buildField(TAG_PAYLOAD_FORMAT_INDICATOR, PAYLOAD_FORMAT_INDICATOR_VALUE),
//		buildField(TAG_POINT_OF_INITIATION_METHOD, initiationMethod),
//		buildField(TAG_MERCHANT_IDENTIFIER_CREDIT_TRANSFER,
//			fmt.Sprintf("%s%s",
//				buildField(TAG_MERCHANT_IDENTIFIER_CREDIT_TRANSFER, aid),
//				buildField(targetType, target))),
//	}
//}

func sanitizePhoneNumber(number, country string) (string, error) {
	switch len(number) {
	case 10:
		if strings.HasPrefix(number, "0") {
			number = strings.TrimPrefix(number, "0")
			number = fmt.Sprintf("00%s%s", country, number)
		} else {
			goto invalidPhoneNum
		}
	case 11:
		if strings.HasPrefix(number, country) {
			number = fmt.Sprintf("00%s", number)
		} else {
			goto invalidPhoneNum
		}
	case 12:
		if strings.HasPrefix(number, fmt.Sprintf("%s0", country)) {
			number = strings.TrimPrefix(number, fmt.Sprintf("%s0", country))
			number = fmt.Sprintf("00%s%s", country, number)
		} else {
			goto invalidPhoneNum
		}
	case 13:
		if !strings.HasPrefix(number, fmt.Sprintf("00%s", country)) {
			goto invalidPhoneNum
		}
	default:
		goto invalidPhoneNum
	}

	return number, nil

invalidPhoneNum:
	return "", ErrInvalidPhoneNumber

}

func isNotInitiationMethod(initiationMethod string) bool {
	return initiationMethod != SINGLE_USE && initiationMethod != MULTIPLE_USE
}

func isNotAID(aid string) bool {
	return aid != MERCHANT_PRESENTED_QR && aid != CUSTOMER_PRESENTED_QR
}

func buildField(id, value string) string {
	return fmt.Sprintf("%s%02d%s", id, len(value), value)
}

func sanitizeTarget(s string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(s, "")
}