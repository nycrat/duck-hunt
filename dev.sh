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

source .env

if [[ $FRONTEND != 1 && $BACKEND != 1 ]]; then
  BOTH=1
fi

if [[ $FRONTEND = 1 || $BOTH = 1 ]]; then
  cd frontend; npm run dev &
  cd ..
fi

if [[ $BACKEND = 1 || $BOTH = 1 ]]; then
  cd backend; go run . $JWT_HS256_KEY $PEPPER $DATABASE_URL_UNPOOLED &
fi

wait
