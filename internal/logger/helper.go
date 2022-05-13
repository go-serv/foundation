package logger

import (
	job "github.com/AgentCoop/go-work"
	"github.com/fatih/color"
	"time"
)

func RegisterStdoutLogger(key interface{}, prewordColor color.Attribute, utfIcon string, on bool) {
	job.RegisterLogger(key, func(args...interface{}) {
		now := time.Now()
		color := color.New(prewordColor, color.Bold)
		color.Printf("[ %s  %s] → ", utfIcon, now.Format(time.StampMilli))
		color.DisableColor()
		fmtStr := args[0].(string)
		if len(args) == 1 {
			color.Println(fmtStr)
		} else {
			color.Printf(fmtStr, args[1:]...)
			color.Println("")
		}
	}, on)
}
