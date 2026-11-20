package tencentcloud

import (
	"encoding/json"
	"fmt"
	"strings"
)

func parseTkeClusterNetworkType(property *string) (string, error) {
	if property == nil || strings.TrimSpace(*property) == "" {
		return "", nil
	}

	var settings struct {
		NetworkType string `json:"NetworkType"`
	}
	if err := json.Unmarshal([]byte(*property), &settings); err != nil {
		return "", fmt.Errorf("parse Cluster.Property network type: %w", err)
	}
	return settings.NetworkType, nil
}
