package components

type ComponentType uint64

const (
	UnknownComponentType  ComponentType = 0 << 0
	IdentityComponentType ComponentType = 1 << 0
)
