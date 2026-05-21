#!/usr/bin/env bash
# cleanup-releases.sh
#
# Keeps a sliding window of the N most recent releases.
# Older releases (and their git tags) are deleted.
#
# Usage:
#   cleanup-releases.sh [keep_count]
#
# Default keep_count = 5.
#
# Required env:
#   GH_TOKEN  - GitHub token with write access to fork releases.

set -euo pipefail

KEEP="${1:-5}"

echo "=== cleanup (keep $KEEP) ==="

victims=$(
  gh release list --limit 200 --json tagName,createdAt \
    --jq '[.[] | select(.tagName | startswith("linux-"))] | sort_by(.createdAt) | reverse | .['"$KEEP"':]  | .[].tagName'
)

if [[ -z "$victims" ]]; then
  echo "  nothing to delete"
  exit 0
fi

while IFS= read -r tag; do
  echo "  deleting $tag"
  gh release delete "$tag" --yes --cleanup-tag
done <<< "$victims"

echo "cleanup complete"
