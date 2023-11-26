package model

import "github.com/mikespook/gorbac"

type RBACManager struct {
	RBAC        *gorbac.RBAC
	Permissions gorbac.Permissions
}

func NewRBACManager(rbac *gorbac.RBAC, permissions gorbac.Permissions) *RBACManager {
	return &RBACManager{
		RBAC:        rbac,
		Permissions: permissions,
	}
}
