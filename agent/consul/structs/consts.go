package structs

const (
	// HealthAny is special, and is used as a wild card,
	// not as a specific state.
	HealthAny      = "any"
	HealthPassing  = "passing"
	HealthWarning  = "warning"
	HealthCritical = "critical"
	HealthMaint    = "maintenance"
)

// KVOp constants give possible operations available in a KVTxn.
type KVOp string

const (
	KVSet            KVOp = "set"
	KVDelete         KVOp = "delete"
	KVDeleteCAS      KVOp = "delete-cas"
	KVDeleteTree     KVOp = "delete-tree"
	KVCAS            KVOp = "cas"
	KVLock           KVOp = "lock"
	KVUnlock         KVOp = "unlock"
	KVGet            KVOp = "get"
	KVGetTree        KVOp = "get-tree"
	KVCheckSession   KVOp = "check-session"
	KVCheckIndex     KVOp = "check-index"
	KVCheckNotExists KVOp = "check-not-exists"
)
