package tusd

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

func (h *UnroutedHandler) log(eventName string, details ...string) {
	LogEvent(h.logger, eventName, details...)
}

func LogEvent(logger *log.Logger, eventName string, details ...string) {
	result := make([]byte, 0, 100)

	result = append(result, `event="`...)
	result = append(result, eventName...)
	result = append(result, `" `...)

	for i := 0; i < len(details); i += 2 {
		result = append(result, details[i]...)
		result = append(result, `="`...)
		result = append(result, details[i+1]...)
		result = append(result, `" `...)
	}

	result = append(result, "\n"...)
	storeLog := logrus.New()
	storeLog.Out = os.Stdout
	file, err := os.OpenFile("tusd-" + time.Now().Format("2006-01-02") + ".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		storeLog.Out = file
	} else {
		storeLog.Info("Failed to log to file, using default stderr")
	}
	storeLog.Info(string(result))
	logger.Output(2, string(result))
}
