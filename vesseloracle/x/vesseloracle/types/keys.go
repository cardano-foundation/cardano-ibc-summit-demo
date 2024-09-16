package types

const (
	// ModuleName defines the module name
	ModuleName = "vesseloracle"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_vesseloracle"
)

var (
	ParamsKey = []byte("p_vesseloracle")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
