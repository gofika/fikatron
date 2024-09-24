#!/bin/bash

# check and update submodule
git submodule update --init --recursive

# check if there are changes
if [[ -n $(git submodule summary) ]]; then
  echo "detected submodule update, commit changes..."
  git config --global user.name 'GitHub Action'
  git config --global user.email 'action@github.com'
  ./generate.sh && git add . && git commit -m "auto update submodule" && git push
else
  echo "no submodule update detected"
fi
