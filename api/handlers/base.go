package handlers

import (
	"github.com/robertsmoto/db_controller_example/repo/models"
)

// BaseHandler will hold everything that controller needs
type BaseHandler struct {
	//userDataLayer models.UserDataLayer
	memberDataLayer models.MemberDataAccessLayer
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(memberDataLayer models.MemberDataAccessLayer) *BaseHandler {
	return &BaseHandler{
		//userDataLayer: userDataLayer,
		memberDataLayer: memberDataLayer,
	}
}
