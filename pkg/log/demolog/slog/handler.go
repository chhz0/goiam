package slog

import (
	"log/slog"
	"os"
	"time"
)

func defaultHandler(lvl *slog.LevelVar) slog.Handler {
	return slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,

		Level: lvl,

		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// 处理自定义的日志级别
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel := level.String()

				switch level {
				case LevelTrace:
					levelLabel = "TRACE"
				}
				a.Value = slog.StringValue(levelLabel)
			}

			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(time.Now().UTC().Format("2006-01-02 15:04:05"))
			}

			return a
		},
	})
}

// Handler 自定义日志后端 slog.Handler
// type Handler struct {
// 	slog.Handler
// }

// // NewHandler 创建新的日志后端 handler
// func NewHandler(w io.Writer, opts *slog.HandlerOptions) *Handler {
// 	return &Handler{
// 		Handler: slog.NewJSONHandler(w, opts),
// 	}
// }

// // Enabled 当前日志级别是否开启
// func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
// 	return h.Handler.Enabled(ctx, level)
// }

// // Handle 处理日志记录，仅在 Enabled() 返回 true 时才会被调用
// func (h *Handler) Handle(ctx context.Context, record slog.Record) error {
// 	record.Add("customlog", "handler")
// 	return h.Handler.Handle(ctx, record)
// }

// // WithAttrs 从现有的 handler 创建一个新的 handler，并将新增属性附加到新的 handler
// func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
// 	return h.Handler.WithAttrs(attrs)
// }

// // WithGroup 从现有的 handler 创建一个新的 handler，并将指定分组附加到新的 handler
// func (h *Handler) WithGroup(name string) slog.Handler {
// 	return h.Handler.WithGroup(name)
// }
