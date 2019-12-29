package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/mdp/qrterminal/v3"
	"go-promptpay-qr"
	"os"
)

const (
	FLAG_PHONE              = "phone"
	FLAG_PHONE_COUNTRY      = "phone_country"
	FLAG_E_WALLET           = "e_wallet"
	FLAG_NATIONAL_ID        = "national_id"
	FLAG_BANK_ACCOUNT       = "account"
	FLAG_MERCHANT_PRESENTED = "merchantqr"
	FLAG_CUSTOMER_PRESENTED = "customerqr"
	FLAG_MULTIPLE_USE       = "multiple"
	FLAG_SINGLE_USE         = "single"
	FLAG_COUNTRY            = "country"
	FLAG_CURRENCY           = "currency"
	FlAG_AMOUNT             = "amount"
)

var (
	ErrInitiationMethodIsRequired             = errors.New("initiation method is required")
	ErrQRTypeIsRequired                       = errors.New("QR type is required")
	ErrCountryCodeIsRequired                  = errors.New("country code is required")
	ErrCurrencyCodeIsRequired                 = errors.New("currency code is required")
	ErrPhoneNumberRequireCountryCodeParameter = errors.New(fmt.Sprintf("phone number country code is required parameter %s", FLAG_PHONE_COUNTRY))
	ErrTargetIsNotProvided                    = errors.New("target is not provided")
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     FLAG_PHONE,
				Usage:    "qr code with phone number",
				Required: false,
			},
			&cli.StringFlag{
				Name:     FLAG_PHONE_COUNTRY,
				Usage:    "qr code with phone country code",
				Required: false,
			},
			&cli.StringFlag{
				Name:     FLAG_E_WALLET,
				Usage:    "qr code with e-wallet",
				Required: false,
			},
			&cli.StringFlag{
				Name:     FLAG_NATIONAL_ID,
				Usage:    "qr code with national id",
				Required: false,
			},
			&cli.StringFlag{
				Name:     FLAG_BANK_ACCOUNT,
				Usage:    "qr code with bank account",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     FLAG_MERCHANT_PRESENTED,
				Usage:    "merchant presented qr code",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     FLAG_CUSTOMER_PRESENTED,
				Usage:    "customer presented qr code",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     FLAG_MULTIPLE_USE,
				Usage:    "multiple use qr code",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     FLAG_SINGLE_USE,
				Usage:    "single use qr code",
				Required: false,
			},
			&cli.StringFlag{
				Name:     FLAG_COUNTRY,
				Usage:    "country code",
				Required: true,
			},
			&cli.StringFlag{
				Name:     FLAG_CURRENCY,
				Usage:    "currency code",
				Required: true,
			},
			&cli.Float64Flag{
				Name:     FlAG_AMOUNT,
				Usage:    "amount",
				Required: false,
			},
		},
		Name:   "go-promptpay-qr",
		Usage:  "generate qr code",
		Action: Action,
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func Action(context *cli.Context) error {
	targetChoose := false
	builder := go_promptpay_qr.NewBuilder()

	if context.Bool(FLAG_SINGLE_USE) {
		builder = builder.WithSingleUse()
	} else if context.Bool(FLAG_MULTIPLE_USE) {
		builder = builder.WithMultipleUse()
	} else {
		return ErrInitiationMethodIsRequired
	}

	if context.Bool(FLAG_MERCHANT_PRESENTED) {
		builder = builder.WithMerchantPresentedQR()
	} else if context.Bool(FLAG_CUSTOMER_PRESENTED) {
		builder = builder.WithCustomerPresentedQR()
	} else {
		return ErrQRTypeIsRequired
	}

	country := context.String(FLAG_COUNTRY)
	if country == "" {
		return ErrCountryCodeIsRequired
	}

	builder = builder.WithCountry(country)

	currency := context.String(FLAG_CURRENCY)
	if country == "" {
		return ErrCurrencyCodeIsRequired
	}

	builder = builder.WithCurrency(currency)

	eWallet := context.String(FLAG_E_WALLET)
	phoneNumber := context.String(FLAG_PHONE)
	phoneCountry := context.String(FLAG_PHONE_COUNTRY)
	nationId := context.String(FLAG_NATIONAL_ID)
	bankAccountId := context.String(FLAG_BANK_ACCOUNT)

	if eWallet != "" {
		b, err := builder.WithEWallet(eWallet)
		if err != nil {
			return err
		}

		targetChoose = true
		builder = b
	} else if phoneNumber != "" {
		if phoneCountry == "" {
			return ErrPhoneNumberRequireCountryCodeParameter
		}

		b, err := builder.WithPhoneNumber(phoneNumber, phoneCountry)
		if err != nil {
			return err
		}

		targetChoose = true
		builder = b
	} else if nationId != "" {
		b, err := builder.WithNationalID(nationId)
		if err != nil {
			return err
		}

		targetChoose = true
		builder = b
	} else if bankAccountId != "" {
		builder = builder.WithBankAccount(bankAccountId)
		targetChoose = true
	}

	if targetChoose != true {
		return ErrTargetIsNotProvided
	}

	if amount := context.Float64(FlAG_AMOUNT); amount != 0 {
		builder = builder.WithAmount(amount)
	}

	s, err := builder.Build()
	if err != nil {
		return err
	}

	config := qrterminal.Config{
		Level: qrterminal.L,
		Writer: os.Stdout,
		BlackChar: qrterminal.WHITE,
		WhiteChar: qrterminal.BLACK,
		QuietZone: 1,
	}

	qrterminal.GenerateWithConfig(s, config)

	return nil
}
