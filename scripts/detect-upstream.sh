#!/usr/bin/env bash
# detect-upstream.sh
#
# Queries tailscale/tailscale for the latest stable release tag (vX.Y.Z) and
# decides whether a new build is needed (based on whether a Release already
# exists for that upstream version).
#
# Outputs are written to $GITHUB_OUTPUT in key=value form.
#
# Required env:
#   GH_TOKEN       - GitHub token with read access to upstream + fork releases
#   FORCE_REBUILD  - (optional) set to "true" to rebuild even if release exists

set -euo pipefail

UPSTREAM_REPO="tailscale/tailscale"
FORCE_REBUILD="${FORCE_REBUILD:-false}"

# --- detect latest stable vX.Y.Z tag ----------------------------------------
# Tailscale uses odd minor versions (1.97.x) for unstable and even (1.96.x) for
# stable. We track even-minor tags only.
latest_tag=$(
  gh api --paginate "repos/${UPSTREAM_REPO}/tags" --jq '.[].name' \
    | grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$' \
    | awk -F. '($2 % 2) == 0' \
    | sort -V \
    | tail -1
)

if [[ -z "$latest_tag" ]]; then
  echo "ERROR: no stable v*.*.* tag found on ${UPSTREAM_REPO}" >&2
  exit 1
fi

# v1.96.4 -> 1.96.4
upstream_version="${latest_tag#v}"

echo "latest upstream stable tag: $latest_tag (version $upstream_version)"

# --- check if we already have a release for this version --------------------
release_tag="linux-v${upstream_version}"
release_exists=false

if gh release view "$release_tag" --json tagName >/dev/null 2>&1; then
  release_exists=true
  echo "release $release_tag already exists"
else
  echo "no release found for $release_tag"
fi

should_build=false
if [[ "$release_exists" == "false" ]] || [[ "$FORCE_REBUILD" == "true" ]]; then
  should_build=true
fi

# --- emit outputs ------------------------------------------------------------
{
  echo "upstream-tag=${latest_tag}"
  echo "upstream-version=${upstream_version}"
  echo "should-build=${should_build}"
  echo "today=$(date -u +%Y-%m-%d)"
} | tee -a "${GITHUB_OUTPUT:-/dev/stdout}"
