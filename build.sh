#!/usr/bin/env bash
set -e
cd vue
yarn build
cd ..
astilectron-bundler -v
