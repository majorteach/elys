package keeper

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elys-network/elys/x/margin/types"
)

func (k Keeper) ForceCloseShort(ctx sdk.Context, mtp *types.MTP, pool *types.Pool, takeFundPayment bool, baseCurrency string) (sdk.Int, error) {
	repayAmount := sdk.ZeroInt()
	// Retrieve AmmPool
	ammPool, err := k.GetAmmPool(ctx, mtp.AmmPoolId, mtp.TradingAsset)
	if err != nil {
		return math.ZeroInt(), err
	}

	err = k.TakeOutCustody(ctx, *mtp, pool, mtp.Custody)
	if err != nil {
		return math.ZeroInt(), err
	}

	// Estimate swap and repay
	repayAmt, err := k.EstimateAndRepay(ctx, *mtp, *pool, ammPool, mtp.Custody, baseCurrency)
	if err != nil {
		return math.ZeroInt(), err
	}

	repayAmount = repayAmount.Add(repayAmt)

	// Hooks after margin position closed
	if k.hooks != nil {
		k.hooks.AfterMarginPositionClosed(ctx, ammPool, *pool)
	}

	return repayAmount, nil
}
