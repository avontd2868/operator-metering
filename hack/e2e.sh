#!/bin/bash
set -e

: "${KUBECONFIG:?}"
: "${METERING_NAMESPACE:=metering-e2e}"

ROOT_DIR=$(dirname "${BASH_SOURCE}")/..
source "${ROOT_DIR}/hack/common.sh"
source "${ROOT_DIR}/hack/lib/tests.sh"

export DEPLOY_SCRIPT="${DEPLOY_SCRIPT:-deploy-e2e.sh}"
export TEST_SCRIPT="$ROOT_DIR/hack/run-e2e-tests.sh"

export TEST_LOG_FILE="${TEST_LOG_FILE:-e2e-tests.log}"
export DEPLOY_LOG_FILE="${DEPLOY_LOG_FILE:-e2e-deploy.log}"
export TEST_TAP_FILE="${TEST_TAP_FILE:-e2e-tests.tap}"

echo "\$KUBECONFIG=$KUBECONFIG"
echo "\$METERING_NAMESPACE=$METERING_NAMESPACE"
echo "\$METERING_OPERATOR_DEPLOY_REPO=$METERING_OPERATOR_DEPLOY_REPO"
echo "\$REPORTING_OPERATOR_DEPLOY_REPO=$REPORTING_OPERATOR_DEPLOY_REPO"
echo "\$METERING_OPERATOR_DEPLOY_TAG=$METERING_OPERATOR_DEPLOY_TAG"
echo "\$REPORTING_OPERATOR_DEPLOY_TAG=$REPORTING_OPERATOR_DEPLOY_TAG"

"$ROOT_DIR/hack/e2e-test-runner.sh"
