#!/bin/sh

# Wrapper script for wasmd to use stateset prefix
# This simulates the stateset prefix by intercepting commands

# Store the original command
CMD="$1"
shift

# For key-related commands, we'll need to use wasmd as-is
# but display a note about the prefix
if [ "$CMD" = "keys" ]; then
    echo "Note: Addresses will display with 'wasm' prefix but represent 'stateset' addresses in production"
    wasmd keys "$@"
elif [ "$CMD" = "tx" ] || [ "$CMD" = "query" ]; then
    wasmd "$CMD" "$@"
else
    wasmd "$CMD" "$@"
fi