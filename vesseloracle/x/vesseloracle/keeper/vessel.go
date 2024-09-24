package keeper

import (
	"context"

	"vesseloracle/x/vesseloracle/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetVessel set a specific vessel in the store from its index
func (k Keeper) SetVessel(ctx context.Context, vessel types.Vessel) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VesselKeyPrefix))
	b := k.cdc.MustMarshal(&vessel)
	store.Set(types.VesselKey(
		vessel.Imo,
		vessel.Ts,
		vessel.Source,
	), b)
}

// GetVessel returns a vessel from its index
func (k Keeper) GetVessel(
	ctx context.Context,
	imo string,
	ts uint64,
	source string,

) (val types.Vessel, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VesselKeyPrefix))

	b := store.Get(types.VesselKey(
		imo,
		ts,
		source,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVessel removes a vessel from the store
func (k Keeper) RemoveVessel(
	ctx context.Context,
	imo string,
	ts uint64,
	source string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VesselKeyPrefix))
	store.Delete(types.VesselKey(
		imo,
		ts,
		source,
	))
}

// GetAllVessel returns all vessel
func (k Keeper) GetAllVessel(ctx context.Context) (list []types.Vessel) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VesselKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Vessel
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
