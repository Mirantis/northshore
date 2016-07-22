#!/bin/bash
#
# http://codeinthehole.com/writing/tips-for-using-a-Git-pre-commit-hook/
#

#
# Build UI distribution if there are some UI files to commit
#
UI_DIR='ui/'
UI_DIST='ui/dist/'
UI_PATTERN='^ui\/.+$'
UI_DIST_STOP=`git diff --name-only | grep -E $UI_PATTERN | wc -l`
UI_DIST_TRIG=`git diff --cached --name-only | grep -E $UI_PATTERN | wc -l`

function on_dist_error {
    git reset HEAD $UI_DIST
    git checkout HEAD -- $UI_DIST
    git status
    echo -e '\n\n\t Commit CANCELED due to distribution problem ';
    echo -e "\t Get more info in ${UI_DIR}npm-debug.log \n\n";
}

function on_dist_success {
    echo -e '\n\t Adding UI distribution to commit \n\n'
    git add $UI_DIST
    git status
}

if [ $UI_DIST_TRIG -gt 0 ]
then
    echo -e "\n\n\t Found some UI files staged to commit.\n\t Going to build the UI distribution. "
    if [ $UI_DIST_STOP -gt 0 ]
    then
        echo -e "\n\t Found changed UI files not staged to commit.\n\t It forbids the UI distribution."
        echo -e "\n\t Stash them or commit with --no-verify "
        exit 2
    fi

    echo -e "\n\t To cancel press CONTROL-C and commit with --no-verify \n\t ENTER to continue "
    # https://stackoverflow.com/a/10015707/4825998
    exec < /dev/tty && read _
    echo -e "\n\t Building the UI distribution \n\n"

    pushd $UI_DIR
    # npm run build || (echo "EX $?"; popd; on_dist_error; exit 1) # braces cause to local exit
    npm run build
    if [ $? -gt 0 ]
    then
        popd
        on_dist_error
        exit 1
    else
        popd
        on_dist_success
        exit 0
    fi
fi
