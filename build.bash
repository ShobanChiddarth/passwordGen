#!/usr/bin/env bash
set -euo pipefail

outdir=build
app="$(basename "$(pwd)")"   # binary name based on current directory
targets=(
  "linux amd64"
  "linux arm64"
  "darwin amd64"
  "darwin arm64"
  "windows amd64"
)

rm -rf "$outdir" && mkdir -p "$outdir"

for t in "${targets[@]}"; do
  read -r GOOS GOARCH <<<"$t"
  ext=""
  [ "$GOOS" = "windows" ] && ext=".exe"
  out="$outdir/${app}-${GOOS}-${GOARCH}${ext}"
  echo "Building $out from $(pwd)..."
  CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w" -o "$out" .
done

touch "$outdir"/.gitkeep
