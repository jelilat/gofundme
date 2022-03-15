package keeper

import (
	"encoding/binary"

	"github.com/cosmonaut/gofundme/x/gofundme/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetGofundmeCount get the total number of gofundme
func (k Keeper) GetGofundmeCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.GofundmeCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetGofundmeCount set the total number of gofundme
func (k Keeper) SetGofundmeCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.GofundmeCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendGofundme appends a gofundme in the store with a new id and update the count
func (k Keeper) AppendGofundme(
	ctx sdk.Context,
	gofundme types.Gofundme,
) uint64 {
	// Create the gofundme
	count := k.GetGofundmeCount(ctx)

	// Set the ID of the appended value
	gofundme.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GofundmeKey))
	appendedValue := k.cdc.MustMarshal(&gofundme)
	store.Set(GetGofundmeIDBytes(gofundme.Id), appendedValue)

	// Update gofundme count
	k.SetGofundmeCount(ctx, count+1)

	return count
}

// SetGofundme set a specific gofundme in the store
func (k Keeper) SetGofundme(ctx sdk.Context, gofundme types.Gofundme) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GofundmeKey))
	b := k.cdc.MustMarshal(&gofundme)
	store.Set(GetGofundmeIDBytes(gofundme.Id), b)
}

// GetGofundme returns a gofundme from its id
func (k Keeper) GetGofundme(ctx sdk.Context, id uint64) (val types.Gofundme, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GofundmeKey))
	b := store.Get(GetGofundmeIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGofundme removes a gofundme from the store
func (k Keeper) RemoveGofundme(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GofundmeKey))
	store.Delete(GetGofundmeIDBytes(id))
}

// GetAllGofundme returns all gofundme
func (k Keeper) GetAllGofundme(ctx sdk.Context) (list []types.Gofundme) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GofundmeKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Gofundme
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetGofundmeIDBytes returns the byte representation of the ID
func GetGofundmeIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetGofundmeIDFromBytes returns ID in uint64 format from a byte array
func GetGofundmeIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
