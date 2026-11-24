pkgname=tailscale-systray-bin
pkgver=1.106.0
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
sha256sums=('20388c17042d01700aed06751ded2191f63f5aed2527840daa1958a04135105c')
options=(!debug !strip)

package() {
    cp -a "$srcdir/usr" "$pkgdir/"
}
