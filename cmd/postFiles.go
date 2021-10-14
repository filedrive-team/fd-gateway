package gateway

import (
	"context"
	"os"
)

func PostFiles(ctx context.Context, path string) (bool, error) {
	file, err := os.Open(path)

	if err != nil {
		return false, nil
	}

	// post the file to estuary node

	return true, nil
}
