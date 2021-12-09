# This is an example PKGBUILD file. Use this as a start to creating your own,
# and remove these comments. For more information, see 'man PKGBUILD'.
# NOTE: Please fill out the license field for your package! If it is unknown,
# then please put 'unknown'.

Maintainer: Barış İnandıoğlu <68742481+baris-inandi@users.noreply.github.com>
pkgname=barley
pkgver=1.0.1
pkgrel=1
epoch=
pkgdesc="A cli that shows a GitHub-like language usage statistics bar."
arch=(x86_64)
url="https://github.com/baris-inandi/barley"
license=('GPLv2')
groups=()
depends=()
makedepends=(git)
checkdepends=()
optdepends=()
provides=(barley)
conflicts=()
replaces=()
backup=()
options=()
install=
changelog=
source=("git+$url")
noextract=()
md5sums=('SKIP')
validpgpkeys=()

pkgver() {
	cd "${_pkgname}"
	printf "1.0.1.r%s.%s" "${git rev-list --count HEAD}" "${git rev-parse --short HEAD}"
}

build() {
	cd barley
	make
}

package() {
	cd "$pkgname-$pkgver"
	make DESTDIR="$pkgdir/" install
}
