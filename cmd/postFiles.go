package gateway

import (
	"context"
	"fmt"
	"os"
)

func PostFiles(ctx context.Context, path string) (bool, error) {
	file, err := os.Open(path)

	if err != nil {
		return false, nil
	}

	fmt.Println(file)

	// post the file to estuary node

	return true, nil
}
