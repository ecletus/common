package common

type WithID interface {
	GetID() string
}

type WithIDSetter interface {
	SetID(value string)
}

type IDSetter interface {
	WithID
	WithIDSetter
}
