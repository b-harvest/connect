package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type OracleHooks interface {
	AfterSanctionListAdded(ctx sdk.Context, sanctionList []SanctionItem) error

	AfterSanctionListRemoved(ctx sdk.Context, sanctionList []SanctionItem) error
}

var _ OracleHooks = &MultiOracleHooks{}

type MultiOracleHooks []OracleHooks

func NewMultiOracleHooks(hooks ...OracleHooks) MultiOracleHooks {
	return hooks
}

func (h MultiOracleHooks) AfterSanctionListAdded(ctx sdk.Context, sanctionList []SanctionItem) error {
	for i := range h {
		if err := h[i].AfterSanctionListAdded(ctx, sanctionList); err != nil {
			return err
		}
	}

	return nil
}

func (h MultiOracleHooks) AfterSanctionListRemoved(ctx sdk.Context, sanctionList []SanctionItem) error {
	for i := range h {
		if err := h[i].AfterSanctionListRemoved(ctx, sanctionList); err != nil {
			return err
		}
	}

	return nil
}
