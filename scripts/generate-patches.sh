#!/usr/bin/env bash
# generate-patches.sh
#
# Extracts the systray-iconname patch from the dev branch as a single
# squashed diff. Only files we own are included (client/systray/ + go.mod).
#
# Usage:
#   generate-patches.sh <base-ref> <dev-branch> <output-dir>
#
# The caller must have fetched both the base ref and the dev branch.

set -euo pipefail

BASE_REF="${1:?Usage: generate-patches.sh <base-ref> <dev-branch> <output-dir>}"
DEV_BRANCH="${2:?dev branch required (e.g. origin/systray-iconname)}"
OUT_DIR="${3:?output directory required}"

mkdir -p "$OUT_DIR"

echo "base ref:   $BASE_REF"
echo "dev branch: $DEV_BRANCH"

git diff "$BASE_REF".."$DEV_BRANCH" -- client/systray/ go.mod \
  > "$OUT_DIR/050-systray-iconname.patch"

patch_size=$(wc -l < "$OUT_DIR/050-systray-iconname.patch")
if [[ "$patch_size" -eq 0 ]]; then
  echo "ERROR: generated patch is empty — no client/systray/ or go.mod changes between $BASE_REF and $DEV_BRANCH" >&2
  exit 1
fi

echo "generated patch ($patch_size lines):"
ls -lh "$OUT_DIR/050-systray-iconname.patch"
