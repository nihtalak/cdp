// Code generated by cdpgen. DO NOT EDIT.

// Package storage implements the Storage domain.
package storage

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the Storage domain.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the Storage domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// ClearDataForOrigin invokes the Storage method. Clears storage for origin.
func (d *domainClient) ClearDataForOrigin(ctx context.Context, args *ClearDataForOriginArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.clearDataForOrigin", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.clearDataForOrigin", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "ClearDataForOrigin", Err: err}
	}
	return
}

// GetUsageAndQuota invokes the Storage method. Returns usage and quota in bytes.
func (d *domainClient) GetUsageAndQuota(ctx context.Context, args *GetUsageAndQuotaArgs) (reply *GetUsageAndQuotaReply, err error) {
	reply = new(GetUsageAndQuotaReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Storage.getUsageAndQuota", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Storage.getUsageAndQuota", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Storage", Op: "GetUsageAndQuota", Err: err}
	}
	return
}
