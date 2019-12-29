# Go Prompt Pay QR

Prompt pay code generator library in go comes with builder 
pattern and client that generates qr code in terminal.

### Install

```
go get github.com/tspn/go-promptpay-qr
```

### How to use
```go
builder, _ := NewBuilder().
    WithMultipleUse().
    WithMerchantPresentedQR().
    WithCountry(THAILAND).
    WithCurrency(THAI_BAHT).
    WithPhoneNumber("0000000000", "66")
code, _ := builder.Build()

//code = 00020101021129370016A000000677010111011300660000000005802TH530376463048956
```

### Build the CLI
```
in  cli folder
$go build -o cli *.go
```
 