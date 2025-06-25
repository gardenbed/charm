// Package ui provides a user interface for command-line applications.
package ui

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// Level represents the verbosity level.
type Level int

const (
	// Trace shows all messages.
	Trace Level = iota
	// Debug shows Debug, Info, Warn, and Error messages.
	Debug
	// Info shows Info, Warn, and Error messages.
	Info
	// Warn shows Warn and Error messages.
	Warn
	// Error shows only Error messages.
	Error
	// None does not show any messages.
	None
)

// UI is the interface for interacting with users in command-line applications.
type UI interface {
	// Output method independent of the verbosity level
	Printf(format string, a ...interface{})

	// Leveled output methods
	GetLevel() Level
	SetLevel(l Level)
	Tracef(s Style, format string, a ...interface{})
	Debugf(s Style, format string, a ...interface{})
	Infof(s Style, format string, a ...interface{})
	Warnf(s Style, format string, a ...interface{})
	Errorf(s Style, format string, a ...interface{})
}

// ui implements the UI interface.
type ui struct {
	sync.Mutex
	level       Level
	writer      io.WriteCloser
	errorWriter io.WriteCloser
}

// New creates a new user interface.
// This is a concurrent-safe UI and can be used across multiple Go routines.
func New(level Level) UI {
	return &ui{
		level:       level,
		writer:      os.Stdout,
		errorWriter: os.Stderr,
	}
}

func (u *ui) Printf(format string, a ...interface{}) {
	u.Lock()
	defer u.Unlock()

	s := fmt.Sprintf(format, a...)
	_, _ = fmt.Fprintln(u.writer, s)
}

func (u *ui) GetLevel() Level {
	u.Lock()
	defer u.Unlock()

	return u.level
}

func (u *ui) SetLevel(l Level) {
	u.Lock()
	defer u.Unlock()

	u.level = l
}

func (u *ui) Tracef(style Style, format string, a ...interface{}) {
	u.Lock()
	defer u.Unlock()

	if u.level <= Trace {
		s := style.sprintf(format, a...)
		_, _ = fmt.Fprintln(u.writer, s)
	}
}

func (u *ui) Debugf(style Style, format string, a ...interface{}) {
	u.Lock()
	defer u.Unlock()

	if u.level <= Debug {
		s := style.sprintf(format, a...)
		_, _ = fmt.Fprintln(u.writer, s)
	}
}

func (u *ui) Infof(style Style, format string, a ...interface{}) {
	u.Lock()
	defer u.Unlock()

	if u.level <= Info {
		s := style.sprintf(format, a...)
		_, _ = fmt.Fprintln(u.writer, s)
	}
}

func (u *ui) Warnf(style Style, format string, a ...interface{}) {
	u.Lock()
	defer u.Unlock()

	if u.level <= Warn {
		s := style.sprintf(format, a...)
		_, _ = fmt.Fprintln(u.writer, s)
	}
}

func (u *ui) Errorf(style Style, format string, a ...interface{}) {
	u.Lock()
	defer u.Unlock()

	if u.level <= Error {
		s := style.sprintf(format, a...)
		_, _ = fmt.Fprintln(u.errorWriter, s)
	}
}
