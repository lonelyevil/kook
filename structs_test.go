package kook

import "testing"

func TestRolePermission_HasPermission(t *testing.T) {
	get := RolePermissionAdmin.HasPermission(RolePermissionAdmin)
	if !get {
		t.Error(get)
	}
	get = RolePermissionAdmin.HasPermission(RolePermissionManageMessage)
	if !get {
		t.Error(get)
	}
}
