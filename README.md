# tailscale-systray — Theme-Native Tray Icons (Fork)

Automated fork of [tailscale/tailscale](https://github.com/tailscale/tailscale) that patches `client/systray/logo.go` to broadcast `SetIconName` over D-Bus instead of `IconPixmap`, so GNOME Shell (and other StatusNotifierItem hosts) resolve tray icons from your installed icon theme (Papirus, Tela, etc.) instead of rendering hardcoded pixels.

## How it works

```
detect-upstream → rebase-patch → build-linux → release → cleanup
```

1. **Detect**: daily check for new `v*.*.*` stable tags on upstream
2. **Rebase**: auto-rebase `systray-iconname` onto the upstream tag
3. **Build**: ubuntu-24.04 + Go — check out the tag directly from `tailscale/tailscale`, apply the patch, `go build ./cmd/systray`
4. **Release**: GitHub release with `tailscale-systray-VERSION-linux-x86_64.tar.gz`, bumps PKGBUILD
5. **Cleanup**: keep 5 most recent releases

The pipeline reads upstream source directly instead of pushing upstream `main`
and release tags into this fork, so upstream-only GitHub workflows are not
triggered here.

## Patches applied

**Theme-native tray icons** — generated at build time from `systray-iconname` branch:
- `client/systray/logo.go` — `systray.SetIconName()` instead of pixel-rendered `SetIcon()`
- `go.mod` — `replace fyne.io/systray => ../systray-fork` (wires in [`nnfewl/systray@set-icon-name`](https://github.com/nnfewl/systray/tree/set-icon-name) which adds `SetIconName` support)

Icon names looked up by the patched binary:
- `tailscale-connected`
- `tailscale-disconnected`
- `tailscale-exit-node-online`
- `tailscale-exit-node-offline`

## Install (Arch — recommended)

```bash
git clone https://github.com/nnfewl/tailscale.git --branch pipeline tailscale-systray-pkg
cd tailscale-systray-pkg
makepkg -si
```

### First-time migration from a hand-rolled unit

If you previously had `~/.config/systemd/user/tailscale-systray.service`:
```bash
systemctl --user disable --now tailscale-systray.service
rm ~/.config/systemd/user/tailscale-systray.service
systemctl --user daemon-reload
systemctl --user enable --now tailscale-systray.service
```

After that, future `makepkg -si` runs only refresh `/usr/bin/tailscale-systray` + the packaged unit. Restart with:
```bash
systemctl --user restart tailscale-systray.service
```

## Install (manual, any distro)

Download `tailscale-systray-VERSION-linux-x86_64.tar.gz` from [Releases](../../releases), then:

```bash
sudo tar -xf tailscale-systray-*-linux-x86_64.tar.gz -C /
systemctl --user daemon-reload
systemctl --user enable --now tailscale-systray.service
```

### Requirements

- An icon theme that ships Tailscale tray icons — Papirus, Tela, etc. (or install the SVGs from [Personal-patch](https://github.com/nnfewl/Personal-patch)'s `tray.sh`)
- Official `tailscale` package (for `tailscaled` daemon — the systray talks to it over LocalAPI)

## Branches

| Branch | Purpose |
|--------|---------|
| `pipeline` | CI/CD scripts, workflows, PKGBUILD (default) |
| `main` | Upstream mirror (not updated by the release pipeline) |
| `systray-iconname` | Dev branch with `client/systray/logo.go` + `go.mod` changes |

## Maintenance

The pipeline auto-rebases `systray-iconname` onto each new upstream tag. If the rebase fails (upstream restructured `client/systray/logo.go`), the workflow opens a GitHub issue with the resolution steps and stops. After resolving locally:

```bash
git fetch origin
git checkout systray-iconname
git rebase vX.Y.Z
# resolve conflicts
git push origin systray-iconname --force-with-lease
gh workflow run release.yml --ref pipeline -f force_rebuild=true
```

## Manual rebuild

```bash
gh workflow run release.yml --ref pipeline -f force_rebuild=true
```
