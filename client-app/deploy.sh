#!/bin/sh
npm run build

rm -rf ../deplay-dashboard/index.html
rm -rf ../deplay-dashboard/static/*
cp -r dist/ ../deploy-dashboard/
cd ../deploy-dashboard
git add .
git commit -am "deploy"
git push origin master
