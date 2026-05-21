// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build cgo || !darwin

package systray

import (
	"bytes"
	"context"
	"sync"
	"time"

	"fyne.io/systray"
)

// SetTheme is a no-op in the SetIconName variant. Theming is delegated to
// the desktop environment's icon theme; the upstream theme flag has no
// effect when icons are looked up by name.
func SetTheme(theme string) {}

type tsLogo struct {
	dots     [9]byte
	iconName string
}

var (
	disconnected = tsLogo{
		dots:     [9]byte{0, 0, 0, 0, 0, 0, 0, 0, 0},
		iconName: "tailscale-disconnected",
	}

	connected = tsLogo{
		dots:     [9]byte{0, 0, 0, 1, 1, 1, 0, 1, 0},
		iconName: "tailscale-connected",
	}

	loading = tsLogo{dots: [9]byte{'l', 'o', 'a', 'd', 'i', 'n', 'g'}}

	loadingLogos = []tsLogo{
		{dots: [9]byte{0, 1, 1, 1, 0, 1, 0, 0, 1}, iconName: "tailscale-connected"},
		{dots: [9]byte{0, 1, 1, 0, 0, 1, 0, 1, 0}, iconName: "tailscale-connected"},
		{dots: [9]byte{0, 1, 1, 0, 0, 0, 0, 0, 1}, iconName: "tailscale-connected"},
		{dots: [9]byte{0, 0, 1, 0, 1, 0, 0, 0, 0}, iconName: "tailscale-connected"},
		{dots: [9]byte{0, 1, 0, 0, 0, 0, 0, 0, 0}, iconName: "tailscale-connected"},
		{dots: [9]byte{0, 0, 0, 0, 0, 1, 0, 0, 0}, iconName: "tailscale-connected"},
		{dots: [9]byte{0, 0, 0, 0, 0, 0, 0, 0, 0}, iconName: "tailscale-connected"},
		{dots: [9]byte{0, 0, 1, 0, 0, 0, 0, 0, 0}, iconName: "tailscale-connected"},
		{dots: [9]byte{0, 0, 0, 0, 0, 0, 1, 0, 0}, iconName: "tailscale-connected"},
		{dots: [9]byte{0, 0, 0, 0, 0, 0, 1, 1, 0}, iconName: "tailscale-connected"},
		{dots: [9]byte{0, 0, 0, 1, 0, 0, 1, 1, 0}, iconName: "tailscale-connected"},
		{dots: [9]byte{0, 0, 0, 1, 1, 0, 0, 1, 0}, iconName: "tailscale-connected"},
		{dots: [9]byte{0, 0, 0, 1, 1, 0, 0, 1, 1}, iconName: "tailscale-connected"},
		{dots: [9]byte{0, 0, 0, 1, 1, 1, 0, 0, 1}, iconName: "tailscale-connected"},
		{dots: [9]byte{0, 1, 0, 0, 1, 1, 1, 0, 1}, iconName: "tailscale-connected"},
	}

	exitNodeOnline = tsLogo{
		dots:     [9]byte{0, 0, 0, 1, 1, 1, 0, 1, 0},
		iconName: "tailscale-exit-node-online",
	}

	exitNodeOffline = tsLogo{
		dots:     [9]byte{0, 0, 0, 1, 1, 1, 0, 1, 0},
		iconName: "tailscale-exit-node-offline",
	}
)

func (logo tsLogo) renderWithBorder(borderUnits int) *bytes.Buffer {
	return bytes.NewBuffer(nil)
}

func setAppIcon(icon tsLogo) {
	if icon.dots == loading.dots {
		startLoadingAnimation()
		return
	}
	stopLoadingAnimation()
	systray.SetIconName(icon.iconName)
}

var (
	loadingMu     sync.Mutex
	loadingCancel func()
)

func startLoadingAnimation() {
	loadingMu.Lock()
	defer loadingMu.Unlock()

	if loadingCancel != nil {
		return
	}

	ctx := context.Background()
	ctx, loadingCancel = context.WithCancel(ctx)

	go func() {
		t := time.NewTicker(500 * time.Millisecond)
		var i int
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				systray.SetIconName(loadingLogos[i].iconName)
				i++
				if i >= len(loadingLogos) {
					i = 0
				}
			}
		}
	}()
}

func stopLoadingAnimation() {
	loadingMu.Lock()
	defer loadingMu.Unlock()

	if loadingCancel != nil {
		loadingCancel()
		loadingCancel = nil
	}
}
