package log

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	log "github.com/nxtcoder17/phuslu-log"
)

type Options struct {
	Prefix string
	Writer io.Writer

	Theme Theme

	LogLevel Level

	SetAsDefaultLogger bool
	DevelopmentMode    bool

	ShowTimestamp bool
	ShowCaller    bool
	ShowLogLevel  *bool
}

type Level uint32

const (
	DebugLevel Level = 0
	InfoLevel  Level = 1
	WarnLevel  Level = 2
	ErrorLevel Level = 3
	// FatalLevel Level = 4
)

type Logger struct {
	logr log.Logger
}

func parseKVs(ze *log.Entry, kv ...any) *log.Entry {
	for i := 0; i < len(kv); i += 2 {
		switch kv[i].(type) {
		case string:
			{
			}
		default:
			{
				fmt.Printf("BAD key (%v), must be string", kv[i])
				continue
			}
		}

		switch tv := kv[i+1].(type) {
		case int:
			ze.Int(kv[i].(string), tv)
		case string:
			ze.Str(kv[i].(string), tv)
		case bool:
			ze.Bool(kv[i].(string), tv)
		default:
			ze.Any(kv[i].(string), tv)
		}
	}
	return ze
}

func (l *Logger) Debug(msg string, kv ...any) {
	parseKVs(l.logr.Debug(), kv...).Msg(msg)
}

func (l *Logger) Info(msg string, kv ...any) {
	parseKVs(l.logr.Info(), kv...).Msg(msg)
}

func (l *Logger) Warn(msg string, kv ...any) {
	parseKVs(l.logr.Warn(), kv...).Msg(msg)
}

func (l *Logger) Error(msg string, kv ...any) {
	parseKVs(l.logr.Error().Caller(1), kv...).Msg(msg)
}

func (l *Logger) Slog() *slog.Logger {
	return l.Slog()
}

func New(opts Options) *Logger {
	if opts.Writer == nil {
		opts.Writer = os.Stderr
	}

	theme := opts.Theme.GetColors()

	fmt.Printf("theme: %+v\n", theme)

	const keyValSeparator = "="

	styledPrefix := func() string {
		if opts.Prefix == "" {
			return ""
		}
		return theme.PrefixStyle.Sprintf(opts.Prefix)
	}()

	writer := func() log.Writer {
		if !opts.DevelopmentMode {
			return log.IOWriter{Writer: opts.Writer}
		}

		return &log.ConsoleWriter{
			ColorOutput:    true,
			QuoteString:    true,
			EndWithMessage: true,
			Writer:         opts.Writer,
			Formatter: func(w io.Writer, args *log.FormatterArgs) (int, error) {
				// return fmt.Fprintf(w, "%s %s\n", styledPrefix, args.Message)

				if styledPrefix != "" {
					fmt.Fprintf(w, "%s ", styledPrefix)
				}

				if opts.ShowLogLevel == nil || *opts.ShowLogLevel {
					fmt.Fprintf(w, "%s ", theme.LogLevelStyles[log.ParseLevel(args.Level)-2].Sprintf(fmt.Sprintf("%-5s", args.Level)))
				}

				fmt.Fprint(w, theme.MessageStyle.Sprintf(args.Message))

				for _, v := range args.KeyValues {
					fmt.Fprintf(w, " %s%s%v", theme.SlogAttrKeyStyle.Sprintf(v.Key), theme.SlogAttrKeyStyle.Sprintf(keyValSeparator), theme.MessageStyle.Sprintf(v.Value))
				}
				return fmt.Fprintf(w, "\n")
			},
		}
	}()

	l := log.Logger{
		Level: func() log.Level {
			return log.Level(opts.LogLevel) + 2
		}(),
		Caller: func() int {
			if opts.ShowCaller {
				return 2
			}
			return 0
		}(),
		ShowTime:     opts.ShowTimestamp,
		TimeField:    "",
		TimeFormat:   "",
		TimeLocation: nil,
		Context:      nil,
		Writer:       writer,
	}

	return &Logger{logr: l}
}
