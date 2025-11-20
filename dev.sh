#!/bin/bash

while getopts "bf" opt; do
  case $opt in
    f)
      FRONTEND=1
      ;;
    b)
      BACKEND=1
      ;;
    \?)
      echo "Invalid option"
      exit 1
      ;;
  esac
done

if [[ $FRONTEND != 1 && $BACKEND != 1 ]]; then
  BOTH=1
fi

if [[ $FRONTEND = 1 || $BOTH = 1 ]]; then
  cd frontend; npm run dev -- --host &
  cd ..
fi

if [[ $BACKEND = 1 || $BOTH = 1 ]]; then
  cd backend; go run cmd/duck-hunt-server/duck-hunt-server.go &
fi

wait
