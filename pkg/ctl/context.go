package ctl

import (
	"context"

	v1 "github.com/pseudomuto/pseudocms/pkg/api/v1"
)

type contextKey string

const (
	adminClientKey  = contextKey("adminClient")
	healthClientKey = contextKey("healthClient")
)

func getAdminClient(ctx context.Context) v1.AdminServiceClient {
	return ctx.Value(adminClientKey).(v1.AdminServiceClient)
}

func setAdminClient(ctx context.Context, c v1.AdminServiceClient) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	return context.WithValue(ctx, adminClientKey, c)
}

func getHealthClient(ctx context.Context) v1.HealthServiceClient {
	return ctx.Value(healthClientKey).(v1.HealthServiceClient)
}

func setHealthClient(ctx context.Context, c v1.HealthServiceClient) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	return context.WithValue(ctx, healthClientKey, c)
}
