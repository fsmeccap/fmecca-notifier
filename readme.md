# Notifier Worker (Go + Pub/Sub + SES) 🚀

Este es un microservicio desarrollado en **Go** diseñado para procesar y despachar notificaciones de correo electrónico de forma asíncrona.

Utiliza **Google Cloud Pub/Sub** como broker de mensajería y **Amazon SES (Simple Email Service)** para el envío confiable de correos mediante plantillas personalizadas.



## 🛠️ Tecnologías utilizadas
* **Go** (Golang)
* **Google Cloud Pub/Sub** (Mensajería asíncrona)
* **Amazon SES v2** (Infraestructura de Email)
* **Slog & Lumberjack** (Estructura de logs y rotación)

## 🏗️ Arquitectura
1. El sistema principal publica un mensaje JSON en un tópico de **Pub/Sub**.
2. Este worker, al estar suscrito, recibe el mensaje automáticamente.
3. El worker mapea los datos a una plantilla de **Amazon SES** y realiza el envío.
4. Se confirma el procesamiento (Ack) para asegurar que ningún correo se pierda.

## 🚀 Configuración
1. Clona el repositorio.
2. Crea un archivo `.env` basado en el `.env.example`.
3. Asegúrate de tener tus credenciales de AWS y GCP configuradas.

```bash
go mod tidy
go run main.go