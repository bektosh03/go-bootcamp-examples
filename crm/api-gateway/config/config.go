package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	StudentServiceAddr  string `envconfig:"STUDENT_SERVICE_ADDR"`
	TeacherServiceAddr  string `envconfig:"TEACHER_SERVICE_ADDR"`
	ScheduleServiceAddr string `envconfig:"SCHEDULE_SERVICE_ADDR"`
	JournalServiceAddr  string `envconfig:"JOURNAL_SERVICE_ADDR"`
	KafkaHost           string `envconfig:"KAFKA_HOST"`
	KafkaPort           string `envconfig:"KAFKA_PORT"`
}

func Load() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
