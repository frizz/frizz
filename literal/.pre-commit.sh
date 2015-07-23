# to install this pre-commit hook, run:
# ln -s ../../.pre-commit.sh .git/hooks/pre-commit

git stash -q --keep-index
go test ./...
RESULT=$?
git stash pop -q
[ $RESULT -ne 0 ] && exit 1
exit 0