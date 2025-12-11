package types

import (
	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type GenesisState struct {
	Params             Params                       `json:"params" yaml:"params"`
	ReserveParams      ReserveParams                `json:"reserve_params" yaml:"reserve_params"`
	Reserve            Reserve                      `json:"reserve" yaml:"reserve"`
	NextVaultId        uint64                       `json:"next_vault_id" yaml:"next_vault_id"`
	Vaults             []Vault                      `json:"vaults" yaml:"vaults"`
	NextDepositId      uint64                       `json:"next_deposit_id" yaml:"next_deposit_id"`
	NextRedemptionId   uint64                       `json:"next_redemption_id" yaml:"next_redemption_id"`
	NextAttestationId  uint64                       `json:"next_attestation_id" yaml:"next_attestation_id"`
	ReserveDeposits    []ReserveDeposit             `json:"reserve_deposits" yaml:"reserve_deposits"`
	RedemptionRequests []RedemptionRequest          `json:"redemption_requests" yaml:"redemption_requests"`
	DailyStats         []DailyMintStats             `json:"daily_stats" yaml:"daily_stats"`
	Attestations       []OffChainReserveAttestation `json:"attestations" yaml:"attestations"`
	ApprovedAttesters  []string                     `json:"approved_attesters" yaml:"approved_attesters"`
}

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:             DefaultParams(),
		ReserveParams:      DefaultReserveParams(),
		Reserve:            Reserve{TotalDeposited: sdk.NewCoins(), TotalValue: sdkmath.ZeroInt(), TotalMinted: sdkmath.ZeroInt()},
		NextVaultId:        1,
		Vaults:             []Vault{},
		NextDepositId:      1,
		NextRedemptionId:   1,
		NextAttestationId:  1,
		ReserveDeposits:    []ReserveDeposit{},
		RedemptionRequests: []RedemptionRequest{},
		DailyStats:         []DailyMintStats{},
		Attestations:       []OffChainReserveAttestation{},
		ApprovedAttesters:  []string{},
	}
}

func (gs GenesisState) Validate() error {
	if err := gs.Params.ValidateBasic(); err != nil {
		return err
	}
	if err := gs.ReserveParams.ValidateBasic(); err != nil {
		return err
	}
	for _, vault := range gs.Vaults {
		if err := vault.ValidateBasic(); err != nil {
			return err
		}
	}
	for _, deposit := range gs.ReserveDeposits {
		if deposit.Id == 0 {
			return ErrReserveDepositNotFound
		}
	}
	for _, redemption := range gs.RedemptionRequests {
		if redemption.Id == 0 {
			return ErrRedemptionNotFound
		}
	}
	for _, stat := range gs.DailyStats {
		if stat.Date == "" {
			return errorsmod.Wrap(ErrInvalidReserve, "daily stat date required")
		}
	}
	for _, att := range gs.Attestations {
		if err := att.ValidateBasic(); err != nil {
			return err
		}
	}
	return nil
}
