# !/bin/bash

########################################################
# Script for integration test
#
# Usage:
#   ./integration-test.sh [basic | dc-static | dc-dynamic]
#
# Description:
#   This script is used to test the functionality of free-ran-ue.
########################################################

DOCKER_PATH='../../docker'
BASIC_COMPOSE_FILE="${DOCKER_PATH}/docker-compose.yaml"
DC_STATIC_COMPOSE_FILE="${DOCKER_PATH}/docker-compose-dc-static.yaml"
DC_DYNAMIC_COMPOSE_FILE="${DOCKER_PATH}/docker-compose-dc-dynamic.yaml"

WEBCONSOLE_BASE_URL='http://127.0.0.1:5000'

WEBCONSOLE_LOGIN_DATA_FILE='webconsole_login_data.json'
WEBCONSOLE_SUBSCRIBER_DATA_FILE='webconsole_subscriber_data.json'

TEST_POOL="basic|dc-static|dc-dynamic"

Usage() {
    echo "Usage: $0 [basic | dc-static | dc-dynamic]"
    exit 1
}

start_docker_compose() {
    if !docker compose -f $1 up -d --wait --wait-timeout 180; then
        echo "Failed to start docker compose!"
        return 1
    fi

    docker ps -a
    return 0
}

stop_docker_compose() {
    if !docker compose -f $1 down; then
        echo "Failed to stop docker compose!"
        return 1
    fi

    return 0
}

webconsole_login() {
    local token=$(curl -s -X POST $WEBCONSOLE_BASE_URL/api/login -H "Content-Type: application/json" -d @$WEBCONSOLE_LOGIN_DATA_FILE | jq -r '.access_token' | xargs)
    if [ -z "$token" ] || [ "$token" = "null" ]; then
        echo "Failed to get token!"
        return 1
    fi

    echo "$token"
    return 0
}

webconsole_subscriber_action() {
    local token=$(webconsole_login)
    if [ -z "$token" ]; then
        echo "Failed to get token!"
        return 1
    fi

    local imsi=$(jq -r '.ueId' "$WEBCONSOLE_SUBSCRIBER_DATA_FILE" | sed 's/imsi-//')
    local plmn_id=$(jq -r '.plmnID' "$WEBCONSOLE_SUBSCRIBER_DATA_FILE")

    case $1 in
        "post")
            if curl -s --fail -X POST $WEBCONSOLE_BASE_URL/api/subscriber/imsi-$imsi/$plmn_id -H "Content-Type: application/json" -H "Token: $token" -d @$WEBCONSOLE_SUBSCRIBER_DATA_FILE; then
                echo "Subscriber created successfully!"
                return 0
            else
                echo "Failed to create subscriber!"
                return 1
            fi
        ;;
        "delete")
            if curl -s --fail -X DELETE $WEBCONSOLE_BASE_URL/api/subscriber/imsi-$imsi/$plmn_id -H "Content-Type: application/json" -H "Token: $token" -d @$WEBCONSOLE_SUBSCRIBER_DATA_FILE; then
                echo "Subscriber deleted successfully!"
                return 0
            else
                echo "Failed to delete subscriber!"
                return 1
            fi
        ;;
    esac
}

main() {
    if [[ ! "$TEST_POOL" =~ "$1" ]]; then
        echo "Invalid test type: $1"
        Usage
        exit 1
    fi

    case $1 in
        "basic")
            start_docker_compose $BASIC_COMPOSE_FILE
            if [ $? -ne 0 ]; then
                echo "Failed to start docker compose!"
                exit 1
            fi

            webconsole_subscriber_action "post"
            if [ $? -ne 0 ]; then
                echo "Failed to create subscriber!"
                stop_docker_compose $BASIC_COMPOSE_FILE
                exit 1
            fi

            webconsole_subscriber_action "delete"
            if [ $? -ne 0 ]; then
                echo "Failed to delete subscriber!"
                stop_docker_compose $BASIC_COMPOSE_FILE
                exit 1
            fi

            stop_docker_compose $BASIC_COMPOSE_FILE
            if [ $? -ne 0 ]; then
                echo "Failed to stop docker compose!"
                exit 1
            fi
        ;;
        "dc-static")
        ;;
        "dc-dynamic")
        ;;
    esac
}

main "$@"