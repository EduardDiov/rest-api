# Create 3 vertically split panes:
# 1: backend golang live reload (using air)
# 2: frontend live reload (using webpack)
# 3: shell

#!/bin/sh                                                                                                   

SESSIONNAME="live"
SCRIPTPATH="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

tmux new-session -s $SESSIONNAME -d
tmux send -t "$SESSIONNAME:" "cd $SCRIPTPATH" Enter
tmux send -t "$SESSIONNAME:" "air" Enter

tmux split-window -v
tmux send -t "$SESSIONNAME:" "cd $SCRIPTPATH/../frontend" Enter
tmux send -t "$SESSIONNAME:" "webpack" Enter

tmux split-window -v
tmux send -t "$SESSIONNAME:" "cd $SCRIPTPATH" Enter
tmux select-layout even-vertical

tmux attach -t $SESSIONNAME
