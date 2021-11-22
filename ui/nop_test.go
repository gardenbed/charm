package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNop(t *testing.T) {
	u := NewNop()

	assert.NotNil(t, u)
}

func TestNopUI_Printf(t *testing.T) {
	u := new(nopUI)
	u.Printf("foo %s", "bar")
}

func TestNopUI_GetLevel(t *testing.T) {
	u := new(nopUI)
	level := u.GetLevel()

	assert.Equal(t, None, level)
}

func TestNopUI_SetLevel(t *testing.T) {
	u := new(nopUI)
	u.SetLevel(None)
}

func TestNopUI_Tracef(t *testing.T) {
	u := new(nopUI)
	u.Tracef(Style{}, "foo %s", "bar")
}

func TestNopUI_Debugf(t *testing.T) {
	u := new(nopUI)
	u.Debugf(Style{}, "foo %s", "bar")
}

func TestNopUI_Infof(t *testing.T) {
	u := new(nopUI)
	u.Infof(Style{}, "foo %s", "bar")
}

func TestNopUI_Warnf(t *testing.T) {
	u := new(nopUI)
	u.Warnf(Style{}, "foo %s", "bar")
}

func TestNopUI_Errorf(t *testing.T) {
	u := new(nopUI)
	u.Errorf(Style{}, "foo %s", "bar")
}
