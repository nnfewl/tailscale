pkgname=tailscale-systray-bin
pkgver=1.100.0
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
sha256sums=('76d87f68e7358bdf4001447739735621387cbe3e191444ac832c41711d1e7728')
options=(!debug !strip)

package() {
    cp -a "$srcdir/usr" "$pkgdir/"
}
