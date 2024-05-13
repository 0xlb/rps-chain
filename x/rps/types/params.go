package types

// DefaultTTL is the default value for the 
// TTL in blocks
const DefaultTTL = 20

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	return Params{
		Ttl: DefaultTTL,
	}
}

// Validate does the sanity check on the params.
func (p Params) Validate() error {
	if p.Ttl == 0 {
		return ErrInvalidTTL
	}
	return nil
}
