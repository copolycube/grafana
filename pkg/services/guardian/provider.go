package guardian

import (
	"context"

	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/services/accesscontrol"
	"github.com/grafana/grafana/pkg/services/sqlstore"
)

type Provider struct{}

func ProvideService(
	store *sqlstore.SQLStore, ac accesscontrol.AccessControl,
	folderPermissionsService accesscontrol.FolderPermissionsService, dashboardPermissionsService accesscontrol.DashboardPermissionsService,
) *Provider {
	if !ac.IsDisabled() {
		// TODO: Fix this hack, see https://github.com/grafana/grafana-enterprise/issues/2935
		InitAccessControlGuardian(store, ac, folderPermissionsService, dashboardPermissionsService)
	} else {
		InitLegacyGuardian(store)
	}
	return &Provider{}
}

func InitLegacyGuardian(store sqlstore.Store) {
	New = func(ctx context.Context, dashId int64, orgId int64, user *models.SignedInUser) DashboardGuardian {
		return newDashboardGuardian(ctx, dashId, orgId, user, store)
	}
}

func InitAccessControlGuardian(
	store sqlstore.Store, ac accesscontrol.AccessControl, folderPermissionsService accesscontrol.FolderPermissionsService,
	dashboardPermissionsService accesscontrol.DashboardPermissionsService,
) {
	New = func(ctx context.Context, dashId int64, orgId int64, user *models.SignedInUser) DashboardGuardian {
		return NewAccessControlDashboardGuardian(ctx, dashId, user, store, ac, folderPermissionsService, dashboardPermissionsService)
	}
}
