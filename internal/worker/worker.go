package worker

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"os"

	"fmecca-notifier/internal/domain"

	"cloud.google.com/go/pubsub"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Worker struct {
	notifier     domain.Notifier
	pubsubClient *pubsub.Client
	subName      string
	logger       *slog.Logger
}

func NewWorker(n domain.Notifier, ps *pubsub.Client, sub string) *Worker {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10,   // Megabytes antes de rotar
		MaxBackups: 2,    // Máximo 2 archivos viejos
		MaxAge:     28,   // Días que se guardan
		Compress:   true, // Comprime los logs viejos (.gz)
	}
	multiWriter := io.MultiWriter(os.Stdout, lumberjackLogger)

	logger := slog.New(slog.NewJSONHandler(multiWriter, nil))

	return &Worker{
		notifier:     n,
		pubsubClient: ps,
		subName:      sub,
		logger:       logger,
	}
}

func (w *Worker) Start(ctx context.Context) {
	sub := w.pubsubClient.Subscription(w.subName)
	w.logger.Info("Worker universal iniciado", "subscription", w.subName)

	err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		defer msg.Ack()

		var req domain.EmailRequest
		if err := json.Unmarshal(msg.Data, &req); err != nil {
			w.logger.Error("Error decodificando JSON universal", "error", err)
			return
		}

		w.logger.Info("Despachando notificación", "to", req.To, "template", req.TemplateName)

		// Enviamos el request tal cual llegó
		err := w.notifier.SendNotification(req)
		if err != nil {
			w.logger.Error("Fallo en Amazon SES", "error", err)
		} else {
			w.logger.Info("✅ Notificación enviada con éxito")
		}
	})

	if err != nil {
		w.logger.Error("Error fatal en el suscriptor", "error", err)
	}
}
