#!/bin/bash
#
# Monitors a directory in the production server, if the directory is not empty
# and contains Go source files it will compile them, stop the running service
# associated to the project, move the compiled binary to the target directory,
# start the new service and delete the deployed source files.
#
# This script is intended to run in a cronjob:
# * * * * * bash /srv/slackbot/autodeploy.sh

DIRECTORY="/srv/slackbot"
INIT_SCRIPT="/etc/init.d/slackbot"
OLD_BINARY="${DIRECTORY}/slackbot.bin"
NEW_BINARY="${DIRECTORY}/slackbot.new"

# The directory does not exists.
if [[ ! -e "$DIRECTORY" ]]; then
	exit 0
fi

# There is no new binary to deploy.
if [[ ! -e "$NEW_BINARY" ]]; then
	exit 0
fi

bash "$INIT_SCRIPT" stop          &>/dev/null
rm -f -- "$OLD_BINARY"            &>/dev/null
chmod +x "$NEW_BINARY"            &>/dev/null
mv -- "$NEW_BINARY" "$OLD_BINARY" &>/dev/null
bash "$INIT_SCRIPT" start         &>/dev/null
