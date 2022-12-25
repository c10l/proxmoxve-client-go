package types

type PVEAPIType interface {
	ToAPIRequestParam() string
}
