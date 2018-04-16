package benchvm

import "errors"

var (
	ErrOutOfAir            = errors.New("BVM Issue! Out of oxygen")
	ErrCodeStoreOutOfAir   = errors.New("BVM Issue! Contract code storage out of oxygen")
	ErrDepth               = errors.New("BVM Issue! Max call depth has been reached.")
	ErrTraceLimitReached   = errors.New("BVM Issue! Log limit has been reached.")
	ErrInsufficientBalance = errors.New("BVM Issue! You don't have a sufficient balance for transfer.")
)
