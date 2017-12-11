#!/bin/bash -

set -eu

VERSION=$(git describe --abbrev=0 --tags)
REVCNT=$(git rev-list --count HEAD)
DEVCNT=$(git rev-list --count $VERSION)
if test $REVCNT != $DEVCNT
then
	VERSION="$VERSION.dev$(expr $REVCNT - $DEVCNT)"
fi
echo "VER: $VERSION"
echo $VERSION > VERSION
git for-each-ref --format="%(refname:short) [%(taggerdate)] [%(subject)] %(body) by:%(authorname)" refs/tags > Changelog
govvv build 
