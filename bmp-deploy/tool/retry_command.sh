#!/bin/bash
set -e
SCRIPT_NAME=${0##*/}
usage(){    
    help="
Usage:  ./${SCRIPT_NAME} Command_id

"
    printf "$help"
}

command_id=$1
[ -n "$command_id" ] || { command_id is null; usage; exit 1; }
[ "$command_id" == "-h" ] && { usage; exit 0; }
cat << EOF | curl -X PUT "127.0.0.1:8801/commands/retryCommand" --data-binary @-
{"offset_command_id":$command_id}
EOF
