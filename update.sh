#!/bin/bash

# check and update submodule
git submodule update --init --recursive
git submodule foreach git pull origin master

# check if there are changes
if [[ -n $(git status -s) ]]; then
  echo "detected submodule update, commit changes..."
  git config --global user.name 'GitHub Action'
  git config --global user.email 'action@github.com'
  sh generate.sh && git commit -m "auto update submodule" && git push
else
  echo "no submodule update detected"
fi
