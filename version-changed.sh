#!/usr/bin/env bash

BRANCH_NAME=${GITHUB_REF##*/}

if [ "${BRANCH_NAME}" = "master" ]; then
  echo "no need check master"
else
  echo "check if version was changed"
  git diff origin/master common.properties | grep +baseVersion
fi
