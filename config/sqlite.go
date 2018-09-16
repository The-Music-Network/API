package config

import "fmt"

type SQLite struct {
	UseMemory bool
	FilePath  string
}

func (this *SQLite) ConnectionString() string {
	if this.UseMemory {
		return ":memory:"
	} else {
		return fmt.Sprintf("file:%s", this.FilePath)
	}
}
