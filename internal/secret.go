package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

type Secrets map[string]string

func loadSecrets(s string) (Secrets, error) {
	secretsFile, err := os.ReadFile(s)
	if err != nil {
		return nil, fmt.Errorf("Read secrets file error: %w\n", err)
	}

	var secrets Secrets
	err = json.Unmarshal(secretsFile, &secrets)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal secrets json error: %w\n", err)
	}

	return secrets, nil
}
