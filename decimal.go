package goshopify

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

// FlexDecimal wraps *decimal.Decimal to handle Shopify API responses that may
// return empty strings ("") for optional decimal fields. The standard
// shopspring/decimal library cannot parse empty strings and returns an error.
type FlexDecimal struct {
	*decimal.Decimal
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// It handles null, empty strings, and valid decimal values.
func (f *FlexDecimal) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == "null" || str == `""` || str == `''` {
		f.Decimal = nil
		return nil
	}

	var d decimal.Decimal
	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}
	f.Decimal = &d
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (f FlexDecimal) MarshalJSON() ([]byte, error) {
	if f.Decimal == nil {
		return []byte("null"), nil
	}
	return json.Marshal(f.Decimal)
}
