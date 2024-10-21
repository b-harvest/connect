package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	marketmaptypes "github.com/skip-mev/connect/v2/x/marketmap/types"
)

// MarketMapKeeper is the expected keeper interface for the market map keeper.
//
//go:generate mockery --name MarketMapKeeper --output ./mocks/ --case underscore
type MarketMapKeeper interface {
	GetMarket(ctx sdk.Context, tickerStr string) (marketmaptypes.Market, error)
}

type OracleHooks interface {
	AfterSanctionListAdded(ctx sdk.Context, sanctionList []SanctionItem) error
	AfterSanctionListRemoved(ctx sdk.Context, sanctionList []SanctionItem) error
}
