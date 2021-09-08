package utils

import (
	"context"
	"os"
	"strconv"
	"time"

	"4d63.com/tz"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

type fn func()

func SetTicker(ctx context.Context, env string, action fn) {
	frequencyStr := os.Getenv(env)
	frequency, err := strconv.Atoi(frequencyStr)
	if frequency == 0 {
		return
	}
	log.Info().Msgf("Config | %s | %ds", env, frequency)
	if err != nil {
		log.Fatal().Err(err).Msgf("Read Env | %s | Fail", env)
	}
	ticker := time.NewTicker(time.Second * time.Duration(frequency))
	for {
		select {
		case <-ticker.C:
			action()
		case <-ctx.Done():
			return
		}
	}
}

func AddSchedule(sc map[string]*cron.Cron, name, timeFormat string, action fn) {
	location, _ := tz.LoadLocation("America/New_York")
	sc[name] = cron.New(
		cron.WithParser(
			cron.NewParser(
				cron.SecondOptional|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dow)),
		cron.WithLocation(location))
	sc[name].AddFunc(timeFormat, func() {
		action()
	})
	sc[name].Start()
}


func MonitorStopCron(ctx context.Context, sc map[string]*cron.Cron) {
	go func() {
		<-ctx.Done()
		for _, v := range sc {
			v.Stop()
		}
	}()
}
