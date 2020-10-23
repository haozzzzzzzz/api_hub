#!/usr/bin/env sh
cd backend
./main 2>&1 &

cd ../frontend_service
npm start