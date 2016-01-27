#!/bin/bash
set -e

gopherjs test -v kego.io/editor/client \
kego.io/editor/client/tree \
kego.io/editor/client/console \
kego.io/editor/shared \
kego.io/editor/shared/connection \
kego.io/editor/shared/messages \
