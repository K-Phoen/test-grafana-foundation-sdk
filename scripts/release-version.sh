#!/usr/bin/env bash

# Exit on error. Append "|| true" if you expect an error.
set -o errexit
# Exit on error inside any functions or subshells.
set -o errtrace
# Do not allow use of undefined vars. Use ${VAR:-} to use an undefined VAR
set -o nounset
# Catch the error in case mysqldump fails (but gzip succeeds) in `mysqldump | gzip`
set -o pipefail

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${__dir}/libs/logs.sh"
source "${__dir}/libs/git.sh"
source "${__dir}/versions.sh"

# These environment variables can be used to alter the behavior of the release script.

DRY_RUN=${DRY_RUN:-"yes"} # Some kind of fail-safe to ensure that we're only pushing something when we mean it.
CLEANUP_WORKSPACE=${CLEANUP_WORKSPACE:-"yes"} # Should the workspace be deleted after the script runs?

GRAFANA_VERSION=${1:-"next"} # version of the schemas/grafana to run the codegen for.
COG_VERSION="v0.0.x" # hardcoded for now

COG_CMD=${COG_CMD:-"cog"} # Command used to run `cog`
GH_CLI_CMD=${GH_CLI_CMD:-"gh"} # Command used to run `gh` (GitHub cli)

KIND_REGISTRY_PATH=${KIND_REGISTRY_PATH:-'../kind-registry'} # Path to the kind-registry
#FOUNDATION_SDK_PATH=${FOUNDATION_SDK_PATH:-'../grafana-foundation-sdk'} # Path to the grafana-foundation-sdk

KIND_REGISTRY_REPO=${KIND_REGISTRY_REPO:-'https://github.com/grafana/kind-registry.git'}
#FOUNDATION_SDK_REPO=${FOUNDATION_SDK_REPO:-'git@github.com:grafana/grafana-foundation-sdk.git'}
FOUNDATION_SDK_REPO=${FOUNDATION_SDK_REPO:-'git@github.com:K-Phoen/test-grafana-foundation-sdk.git'}

#################
### Usage ###
#################

# LOG_LEVEL=7 ./scripts/release-version.sh v10.2.x

#################
### Utilities ###
#################

function clone_kind_registry() {
  local clone_into_dir="${1}"
  shift

  git clone --depth 1 "${KIND_REGISTRY_REPO}" "${clone_into_dir}"
}

function clone_foundation_sdk() {
  local clone_into_dir="${1}"
  shift

  git clone "${FOUNDATION_SDK_REPO}" "${clone_into_dir}"
}

function run_codegen() {
  local extra_parameters="${1}"
  shift

  $COG_CMD generate \
    --config "${__dir}/../.cog/config.yaml" \
    --parameters "${extra_parameters}"
}

function gh_run() (
  local repo_dir=${1}
  shift

  cd "$repo_dir"

  $GH_CLI_CMD "$@"
)

function run_when_safe() {
  local command=${1}
  shift

  if [ "${DRY_RUN}" == "no" ]; then
    ${command} "$@"
  else
    warning "Dry run enabled: skipping execution of \"${command} $*\""
    info "Run this script with DRY_RUN=no to disable dry-run mode."
  fi
}

############
### Main ###
############

workspace_path='./workspace'
codegen_output_path="${workspace_path}/codegen"
foundation_sdk_path="${workspace_path}/foundation-sdk"

function cleanup() {
  debug "Cleaning up workspace"
  rm -rf "${workspace_path}"
}
if [ "${CLEANUP_WORKSPACE}" == "yes" ]; then
  trap cleanup EXIT # run the cleanup() function on exit
fi

release_branch="${GRAFANA_VERSION}+cog-${COG_VERSION}"
build_timestamp="$(date '+%s')"
pr_branch="${release_branch}-preview-${build_timestamp}"

if [ "${DRY_RUN}" == "no" ]; then
  warning "Dry-run is OFF."
else
  notice "Dry-run is ON."
fi

notice "Grafana version: ${GRAFANA_VERSION}"
notice "Cog version: ${COG_VERSION}"
notice "Release branch in grafana-foundation-sdk: ${release_branch}"
debug "kind-registry path: ${KIND_REGISTRY_PATH}"
debug "workspace path: ${workspace_path}"

# Just in case there are leftovers from a previous run.
rm -rf "${workspace_path}"

if [ ! -d "${KIND_REGISTRY_PATH}" ]; then
  info "Cloning kind-registry into ${KIND_REGISTRY_PATH}"
  clone_kind_registry "${KIND_REGISTRY_PATH}"
fi

info "Cloning grafana-foundation-sdk into ${foundation_sdk_path}"
clone_foundation_sdk "${foundation_sdk_path}"

info "Pulling kind-registry@main"
git_run "${KIND_REGISTRY_PATH}" checkout main
git_run "${KIND_REGISTRY_PATH}" pull --ff-only origin main

info "Running cog"
grafana_version_or_main=${GRAFANA_VERSION}
if [ "${grafana_version_or_main}" == "next" ]; then
  grafana_version_or_main="main"
fi

run_codegen "output_dir=${codegen_output_path}/%l,kind_registry_path=${KIND_REGISTRY_PATH},kind_registry_version=${GRAFANA_VERSION},grafana_version=${grafana_version_or_main},cog_version=${COG_VERSION},all_grafana_versions=${ALL_GRAFANA_VERSIONS},release_branch=${release_branch},build_timestamp=${build_timestamp}"

release_branch_exists=$(git_has_branch "${foundation_sdk_path}" "${release_branch}")
if [ "$release_branch_exists" != "0" ]; then
  info "Creating new release branch in grafana-foundation-sdk";
  git_create_orphan_branch "${foundation_sdk_path}" "${release_branch}";
  git_run "${foundation_sdk_path}" commit --allow-empty -m "Initialize release branch for Grafana ${GRAFANA_VERSION} and Cog ${COG_VERSION}"

  info "Pushing release branch ${release_branch}"
  run_when_safe git_run "${foundation_sdk_path}" push origin "${release_branch}"
else
  info "Checking out existing release branch in grafana-foundation-sdk"
  git_run "${foundation_sdk_path}" checkout "${release_branch}"
  git_run "${foundation_sdk_path}" pull --ff-only origin "${release_branch}"
fi

info "Creating release preview branch in grafana-foundation-sdk"
git_run "${foundation_sdk_path}" checkout -b "${pr_branch}"

debug "Copying generated content to grafana-foundation-sdk"
find "${codegen_output_path}" -maxdepth 1 -mindepth 1 -print | while read -r target; do
  target=${target#"$codegen_output_path/"}

  # By removing the language folder before copying the generated output, we make
  # sure that files that might have been generated by a previous release but
  # aren't in the current workspace are pruned.
  rm -rf "${foundation_sdk_path:?}/${target}"
  cp -R "${codegen_output_path}/${target}" "${foundation_sdk_path}"
done

debug "Adding changes to git staging area"
git_run "${foundation_sdk_path}" add .

has_changes=$(git_has_changes "${foundation_sdk_path}")
if [ "${has_changes}" != "0" ]; then
  warning "No changes detected in grafana-foundation-sdk: aborting release."
  exit 0
fi

git_run "${foundation_sdk_path}" commit -m "Automated release"

info "Pushing PR branch ${pr_branch}"
run_when_safe git_run "${foundation_sdk_path}" push origin "${pr_branch}"

debug "Ensuring that ${FOUNDATION_SDK_REPO} will be used by gh"
run_when_safe gh_run "${foundation_sdk_path}" repo set-default "${FOUNDATION_SDK_REPO}"

info "Opening release Pull Request"
run_when_safe gh_run "${foundation_sdk_path}" pr create \
  --base "${release_branch}" \
  --title "Automated release for Grafana ${GRAFANA_VERSION} and Cog ${COG_VERSION}" \
  --body "Automated release."

if [ "${DRY_RUN}" != "no" ]; then
  notice "Review the changes on the ${pr_branch} branch in ${foundation_sdk_path} and re-run this script with DRY_RUN=no to disable dry-run mode."
  notice "Tip: git show HEAD"
fi
