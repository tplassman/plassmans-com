#!/bin/bash

# The purpose of this script is to deploy files to remote enviroinments.
# The remote enironment can be either staging or production.
#
# This script requires a sibling file named ".envsh" that contains named
# variabels as follows:
#
# |_HOME_PATH
# |_APP_DIR
# |_STATIC_DIR
# |_SSH_HOST
# |_STAG_HOME_PATH
# |_PROD_HOME_PATH

# Import env variables
source .env

# Initialize variables
REMOTE_HOME_PATH=""

################################################################################
# Error checks
################################################################################

# Check for required arguments
if [[ -z "$1" || ( "$1" != "staging" && "$1" != "production" ) ]]; then
  echo "Please provide a valid remote environment argument"

  exit
fi

# Set remote environment
if [ "$1" = "staging" ]; then
  REMOTE_HOME_PATH="$STAG_HOME_PATH"
elif [ "$1" = "production" ]; then
  REMOTE_HOME_PATH="$PROD_HOME_PATH"
fi

################################################################################
# File sync
################################################################################

# Build, sync, and restart binary
GOOS=linux GOARCH=amd64 go build -o "$HOME_PATH/$APP_DIR/cmd/server" -v "$HOME_PATH/$APP_DIR/cmd/server.go"
rsync -azh "$HOME_PATH/$APP_DIR/" "$SSH_HOST:$REMOTE_HOME_PATH/$APP_DIR"
ssh "$SSH_HOST" "supervisorctl restart website"
# Sync static files
rsync -azh --filter=":- .gitignore" "$HOME_PATH/$STATIC_DIR/" "$SSH_HOST:$REMOTE_HOME_PATH/$STATIC_DIR"
