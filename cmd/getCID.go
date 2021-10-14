package gateway

import (
	"context"

	"github.com/application-research/estuary/node"
)

type SysFile struct {
}

// The base url
const BASE_URL = ""

func GetByCID(ctx context.Context, targetPath string) (*SysFile, error) {
	/*
	 * The auth token should be passed to the function so that it can call estuary APIs with bear auth
	 * Then try to get from estuary node
	 * and return the file info
	 */
	n, err := node.Setup(ctx, da)

	if err != nil {
		return nil, err
	}

	if n != nil {

	}

	return nil, nil
}
