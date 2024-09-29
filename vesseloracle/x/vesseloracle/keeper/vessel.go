package keeper

import (
	"context"
	"fmt"
	"sort"

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

	k.AddVesselKeyToIndexImo(ctx, types.VesselIndexImo_Key{
		Imo:    vessel.Imo,
		Ts:     vessel.Ts,
		Source: vessel.Source,
	})
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

	k.RemoveVesselKeyFromIndexImo(ctx, types.VesselIndexImo_Key{
		Imo:    imo,
		Ts:     ts,
		Source: source,
	})
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

func (k Keeper) GetVesselsInWindow(ctx context.Context, imo string, intervalWidth uint64, maxItemsCount int32) (vessels []types.Vessel) {
	// check if maxItemsCount is at least 1 otherwise immediately return
	if maxItemsCount == 0 {
		return vessels
	}

	// fetch maxItemsCount with imo and ordered by ts decreasing, then filter interWidth and return
	keyIndex, found := k.GetVesselKeysFromIndexImo(ctx, imo)
	if found == false || len(keyIndex.Keys) == 0 {
		k.Logger().Error("Could not find any vessels for IMO", imo)
		return vessels
	}

	// order keys by TS decreasing
	sort.Slice(keyIndex.Keys, func(i, j int) bool {
		return keyIndex.Keys[i].Ts > keyIndex.Keys[j].Ts
	})
	k.Logger().Info(fmt.Sprint("Keys ordered:", keyIndex.Keys))

	var maxTs = keyIndex.Keys[0].Ts

	// only pick the items within the time window [maxTs - intervalWidth, maxTs]
	intervalWidth = min(intervalWidth, maxTs)
	for keyIdx, key := range keyIndex.Keys {
		if (key.Ts < (maxTs - intervalWidth)) || (keyIdx >= int(maxItemsCount)) {
			keyIndex.Keys = keyIndex.Keys[:keyIdx]
			break
		}
	}

	// fetch items
	for _, key := range keyIndex.Keys {
		vessel, found := k.GetVessel(ctx, key.Imo, key.Ts, key.Source)
		if found {
			vessels = append(vessels, vessel)
		} else {
			k.Logger().Error(fmt.Sprintf("Could not find index %v. Should not happen.", key))
		}
	}

	return vessels
}
