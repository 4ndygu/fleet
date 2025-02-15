package service

import (
	"context"
	"time"

	"github.com/kolide/fleet/server/kolide"
)

func (mw loggingMiddleware) ListHosts(ctx context.Context, opt kolide.ListOptions) ([]*kolide.Host, error) {
	var (
		hosts []*kolide.Host
		err   error
	)

	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "ListHosts",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	hosts, err = mw.Service.ListHosts(ctx, opt)
	return hosts, err
}

func (mw loggingMiddleware) ListHostsPaginated(ctx context.Context, opt kolide.ListOptions) ([]*kolide.Host, error) {
	var (
		hosts []*kolide.Host
		err   error
	)

	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "ListHostsPaginated",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	hosts, err = mw.Service.ListHostsPaginated(ctx, opt)
	return hosts, err
}

func (mw loggingMiddleware) GetHost(ctx context.Context, id uint) (*kolide.Host, error) {
	var (
		host *kolide.Host
		err  error
	)

	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "GetHost",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	host, err = mw.Service.GetHost(ctx, id)
	return host, err
}

func (mw loggingMiddleware) GetHostSummary(ctx context.Context) (*kolide.HostSummary, error) {
	var (
		summary *kolide.HostSummary
		err     error
	)

	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "GetHostSummary",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	summary, err = mw.Service.GetHostSummary(ctx)
	return summary, err
}

func (mw loggingMiddleware) DeleteHost(ctx context.Context, id uint) error {
	var (
		err error
	)

	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "DeleteHost",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.Service.DeleteHost(ctx, id)
	return err
}
