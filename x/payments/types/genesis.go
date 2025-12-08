package types

type GenesisState struct {
	NextPaymentId uint64          `json:"next_payment_id" yaml:"next_payment_id"`
	Payments      []PaymentIntent `json:"payments" yaml:"payments"`
}

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		NextPaymentId: 1,
		Payments:      []PaymentIntent{},
	}
}

func (gs GenesisState) Validate() error {
	for _, payment := range gs.Payments {
		if err := payment.ValidateBasic(); err != nil {
			return err
		}
	}
	return nil
}
