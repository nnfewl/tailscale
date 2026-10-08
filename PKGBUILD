pkgname=tailscale-systray-bin
pkgver=1.104.1
pkgrel=1
pkgdesc='Tailscale system tray with theme-native icons (Papirus/Tela support)'
arch=('x86_64')
url='https://github.com/nnfewl/tailscale'
license=('BSD-3-Clause')
provides=("tailscale-systray=${pkgver}")
conflicts=('tailscale-systray-git')
depends=('glibc' 'tailscale' 'hicolor-icon-theme')
optdepends=('papirus-icon-theme: ships Tailscale tray icons'
            'tela-icon-theme: ships Tailscale tray icons')
source=("https://github.com/nnfewl/tailscale/releases/download/linux-v${pkgver}/tailscale-systray-${pkgver}-linux-x86_64.tar.gz")
sha256sums=('d040bbbcbb70b5b3f7cd5134960c0ecee6c7e5d7187068241ab5b64c839f07d3')
options=(!debug !strip)

package() {
    cp -a "$srcdir/usr" "$pkgdir/"
}
