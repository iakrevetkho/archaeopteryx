package config

import (
	"embed"
	"time"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Version string `default:"" env:"VERSION"`

	Port uint64 `default:"8080" env:"PORT"`

	Log struct {
		Level    logrus.Level `default:"info" env:"LOG_LEVEL"`
		Filename string       `default:"/var/log/archaeopteryx/archaeopteryx.log" env:"LOG_FILE_NAME"`
		// If log file size is bigger than this threshold, it will be rotated
		MaxSizeInMb int `default:"20" env:"LOG_MAX_SIZE_IN_MB"`
		// If old log file has age more than 30 days, it will be deleted
		MaxAgeInDays int  `default:"30" env:"LOG_MAX_AGE_IN_DAYS"`
		MaxBackups   int  `default:"30" env:"LOG_MAX_BACKUPS"`
		Compress     bool `default:"true" env:"LOG_COMPRESS_OLD_FILES"`
	}

	Health struct {
		// Period for sending healthcheck status on Watch method
		WatchUpdatePeriod time.Duration `default:"15s" env:"HEALTH_WATCH_UPDATE_PERIOD"`
	}

	Docs struct {
		// Embed file system with service Swagger documentation
		DocsFS *embed.FS
		// Folder name in FS with swagger docs
		DocsRootFolder string
	}
}
