package keeper_test

import (
	"testing"
	"time"

	appparams "github.com/OmniFlix/omniflixhub/v2/app/params"
	globalfeekeeper "github.com/OmniFlix/omniflixhub/v2/x/globalfee/keeper"
	"github.com/OmniFlix/omniflixhub/v2/x/globalfee/types"
	dbm "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupTestStore(t *testing.T) (sdk.Context, appparams.EncodingConfig, globalfeekeeper.Keeper) {
	t.Helper()
	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	encCfg := appparams.MakeEncodingConfig()
	keyParams := sdk.NewKVStoreKey(types.StoreKey)
	ms.MountStoreWithDB(keyParams, storetypes.StoreTypeIAVL, db)
	require.NoError(t, ms.LoadLatestVersion())

	globalFeeKeeper := globalfeekeeper.NewKeeper(encCfg.Marshaler, keyParams, "omniflix1llyd96levrglxhw6sczgk9wn48t64zkhv4fq0r")

	ctx := sdk.NewContext(ms, tmproto.Header{
		Height:  1234567,
		Time:    time.Date(2020, time.April, 22, 12, 0, 0, 0, time.UTC),
		ChainID: "omniflixhub-test",
	}, false, log.NewNopLogger())

	return ctx, encCfg, globalFeeKeeper
}