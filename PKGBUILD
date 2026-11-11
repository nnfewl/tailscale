pkgname=tailscale-systray-bin
pkgver=1.104.6
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
sha256sums=('e82940756c4ef930c82b6da46442623104f2e53620d64186def7448d74907440')
options=(!debug !strip)

package() {
    cp -a "$srcdir/usr" "$pkgdir/"
}
