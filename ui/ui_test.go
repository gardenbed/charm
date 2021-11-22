package ui

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

type nopCloser struct {
	io.Writer
}

func (w *nopCloser) Close() error {
	return nil
}

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		level Level
	}{
		{
			name:  "Info",
			level: Info,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			u := New(tc.level)
			assert.NotNil(t, u)

			uu, ok := u.(*ui)
			assert.True(t, ok)
			assert.Equal(t, tc.level, uu.level)
			assert.NotNil(t, uu.writer)
			assert.NotNil(t, uu.errorWriter)
		})
	}
}

func TestUI_Printf(t *testing.T) {
	tests := []struct {
		name           string
		format         string
		args           []interface{}
		expectedOutput string
	}{
		{
			name:           "OK",
			format:         "Hello, %s!",
			args:           []interface{}{"World"},
			expectedOutput: "Hello, World!\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			u := &ui{
				writer: &nopCloser{
					Writer: buf,
				},
			}

			u.Printf(tc.format, tc.args...)

			out, err := io.ReadAll(buf)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedOutput, string(out))
		})
	}
}

func TestUI_GetLevel(t *testing.T) {
	tests := []struct {
		name  string
		level Level
	}{
		{
			name:  "Debug",
			level: Debug,
		},
		{
			name:  "Info",
			level: Info,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			u := &ui{
				level: tc.level,
			}

			level := u.GetLevel()

			assert.Equal(t, tc.level, level)
		})
	}
}

func TestUI_SetLevel(t *testing.T) {
	tests := []struct {
		name  string
		level Level
	}{
		{
			name:  "Debug",
			level: Debug,
		},
		{
			name:  "Info",
			level: Info,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			u := &ui{}

			u.SetLevel(tc.level)

			assert.Equal(t, tc.level, u.level)
		})
	}
}

func TestUI_Tracef(t *testing.T) {
	tests := []struct {
		name           string
		style          Style
		format         string
		args           []interface{}
		expectedOutput string
	}{
		{
			name:           "OK_Blue",
			style:          Blue,
			format:         "Hello, %s!",
			args:           []interface{}{"World"},
			expectedOutput: "\x1b[34mHello, World!\x1b[0m\n",
		},
		{
			name:           "OK_Mix",
			style:          Style{BgYellow, FgMagenta, Bold, BlinkSlow},
			format:         "Hello, %s!",
			args:           []interface{}{"World"},
			expectedOutput: "\x1b[43;35;1;5mHello, World!\x1b[0m\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			u := &ui{
				writer: &nopCloser{
					Writer: buf,
				},
			}

			u.Tracef(tc.style, tc.format, tc.args...)

			out, err := io.ReadAll(buf)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedOutput, string(out))
		})
	}
}

func TestUI_Debugf(t *testing.T) {
	tests := []struct {
		name           string
		style          Style
		format         string
		args           []interface{}
		expectedOutput string
	}{
		{
			name:           "OK_Cyan",
			style:          Cyan,
			format:         "Hello, %s!",
			args:           []interface{}{"World"},
			expectedOutput: "\x1b[36mHello, World!\x1b[0m\n",
		},
		{
			name:           "OK_Mix",
			style:          Style{BgYellow, FgMagenta, Bold, BlinkSlow},
			format:         "Hello, %s!",
			args:           []interface{}{"World"},
			expectedOutput: "\x1b[43;35;1;5mHello, World!\x1b[0m\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			u := &ui{
				writer: &nopCloser{
					Writer: buf,
				},
			}

			u.Debugf(tc.style, tc.format, tc.args...)

			out, err := io.ReadAll(buf)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedOutput, string(out))
		})
	}
}

func TestUI_Infof(t *testing.T) {
	tests := []struct {
		name           string
		style          Style
		format         string
		args           []interface{}
		expectedOutput string
	}{
		{
			name:           "OK_Green",
			style:          Green,
			format:         "Hello, %s!",
			args:           []interface{}{"World"},
			expectedOutput: "\x1b[32mHello, World!\x1b[0m\n",
		},
		{
			name:           "OK_Mix",
			style:          Style{BgYellow, FgMagenta, Bold, BlinkSlow},
			format:         "Hello, %s!",
			args:           []interface{}{"World"},
			expectedOutput: "\x1b[43;35;1;5mHello, World!\x1b[0m\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			u := &ui{
				writer: &nopCloser{
					Writer: buf,
				},
			}

			u.Infof(tc.style, tc.format, tc.args...)

			out, err := io.ReadAll(buf)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedOutput, string(out))
		})
	}
}

func TestUI_Warnf(t *testing.T) {
	tests := []struct {
		name           string
		style          Style
		format         string
		args           []interface{}
		expectedOutput string
	}{
		{
			name:           "OK_Yellow",
			style:          Yellow,
			format:         "Hello, %s!",
			args:           []interface{}{"World"},
			expectedOutput: "\x1b[33mHello, World!\x1b[0m\n",
		},
		{
			name:           "OK_Mix",
			style:          Style{BgYellow, FgMagenta, Bold, BlinkSlow},
			format:         "Hello, %s!",
			args:           []interface{}{"World"},
			expectedOutput: "\x1b[43;35;1;5mHello, World!\x1b[0m\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			u := &ui{
				writer: &nopCloser{
					Writer: buf,
				},
			}

			u.Warnf(tc.style, tc.format, tc.args...)

			out, err := io.ReadAll(buf)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedOutput, string(out))
		})
	}
}

func TestUI_Errorf(t *testing.T) {
	tests := []struct {
		name           string
		style          Style
		format         string
		args           []interface{}
		expectedOutput string
	}{
		{
			name:           "OK_Red",
			style:          Red,
			format:         "Hello, %s!",
			args:           []interface{}{"World"},
			expectedOutput: "\x1b[31mHello, World!\x1b[0m\n",
		},
		{
			name:           "OK_Mix",
			style:          Style{BgYellow, FgMagenta, Bold, BlinkSlow},
			format:         "Hello, %s!",
			args:           []interface{}{"World"},
			expectedOutput: "\x1b[43;35;1;5mHello, World!\x1b[0m\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			u := &ui{
				errorWriter: &nopCloser{
					Writer: buf,
				},
			}

			u.Errorf(tc.style, tc.format, tc.args...)

			out, err := io.ReadAll(buf)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedOutput, string(out))
		})
	}
}
