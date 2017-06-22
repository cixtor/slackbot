#!/bin/bash
### BEGIN INIT INFO
# Provides:
# Required-Start:    $remote_fs $syslog
# Required-Stop:     $remote_fs $syslog
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: Start daemon at boot time
# Description:       Enable service provided by daemon.
### END INIT INFO

USERNAME="root"
DIRECTORY="/srv/slackbot"
PROGRAM=$(basename "$0")
COMMAND="${DIRECTORY}/${PROGRAM}.bin"
PID_FILE="${DIRECTORY}/${PROGRAM}.pid"
STDOUT_LOG="${DIRECTORY}/${PROGRAM}.log"

get_pid() {
    cat "$PID_FILE"
}

is_running() {
    [ -f "$PID_FILE" ] && ps "$(get_pid)" > /dev/null 2>&1
}

case "$1" in
    start)
    if is_running; then
        echo "Already started"
    else
        echo "Starting $PROGRAM"
        cd "$DIRECTORY" || exit
        sudo -E -u "$USERNAME" "$COMMAND" 1>> "$STDOUT_LOG" 2>&1 &
        echo $! > "$PID_FILE"
        if ! is_running; then
            echo "Unable to start, see ${STDOUT_LOG}"
            exit 1
        fi
    fi
    ;;
    stop)
    if is_running; then
        printf "Stopping %s.." "$PROGRAM"
        kill "$(get_pid)"
        for _ in {1..10}; do
            if ! is_running; then break; fi
            printf "."
            sleep 1
        done
        echo

        if is_running; then
            echo "Not stopped; may still be shutting down or shutdown may have failed"
            exit 1
        else
            echo "Stopped"
            if [ -f "$PID_FILE" ]; then
                rm -- "$PID_FILE" 2> /dev/null
            fi
        fi
    else
        echo "Not running"
    fi
    ;;
    restart)
    $0 stop
    if is_running; then
        echo "Unable to stop, will not attempt to start"
        exit 1
    fi
    $0 start
    ;;
    status)
    if is_running; then
        echo "Running"
    else
        echo "Stopped"
        exit 1
    fi
    ;;
    *)
    echo "Usage: $0 {start|stop|restart|status}"
    exit 1
    ;;
esac

exit 0
