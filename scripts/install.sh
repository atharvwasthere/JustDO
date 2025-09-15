#!/usr/bin/env bash
set -euo pipefail

REPO="atharvwasthere/JustDO"
OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
case "$ARCH" in
  x86_64) ARCH=amd64 ;;
  aarch64) ARCH=arm64 ;;
  arm64) ARCH=arm64 ;;
  *) echo "Unsupported arch: $ARCH" >&2; exit 1 ;;
esac

TAG="${1:-latest}"
if [ "$TAG" = "latest" ]; then
  TAG=$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" | grep -Po '"tag_name":\s*"\K[^"]+')
fi

NAME="justdo_${TAG#v}_${OS}_${ARCH}"
EXT="tar.gz"
[ "$OS" = "windows" ] && EXT="zip"

URL="https://github.com/$REPO/releases/download/$TAG/${NAME}.${EXT}"
TMP=$(mktemp -d)

echo "Downloading $URL"
curl -fsSL "$URL" -o "$TMP/justdo.${EXT}"

if [ "$EXT" = "zip" ]; then
  unzip -q "$TMP/justdo.${EXT}" -d "$TMP"
else
  tar -xzf "$TMP/justdo.${EXT}" -C "$TMP"
fi

install "$TMP/justdo" /usr/local/bin/justdo
echo "Installed: $(justdo --version)"
