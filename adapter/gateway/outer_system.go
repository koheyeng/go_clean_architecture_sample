package gateway

import (
	"github.com/koheyeng/go_clean_architecture_sample/usecase/users"
)

type outerSystemGateway struct {
}

func NewOuterSystemGateway() users.OuterSystemGateway {
	return &outerSystemGateway{}
}

func (u *outerSystemGateway) Send(userID int) error {

	// Do Something

	return nil
}
