package gateway

import (
	"github.com/koheyeng/go_clean_architecture_sample/usecase"
)

type outerSystemGateway struct {
}

func NewOuterSystemGateway() usecase.OuterSystemGateway {
	return &outerSystemGateway{}
}

func (u *outerSystemGateway) Send(userID int) error {

	// Do Something

	return nil
}
