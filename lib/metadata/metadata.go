package metadata

import (
	"encoding/json"
	"os"
	"time"
)

type T struct {
	BackupHost string `json:"backup_host"`
	Kind       string `json:"kind"`
	FetchTime  time.Duration `json:"fetch_duration"`
}

func (t T) WriteFor(path string) error {
	f, err := os.Create(path + ".scbmeta")
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(t)
}
