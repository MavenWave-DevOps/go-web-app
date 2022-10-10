#!/bin/zsh
PROJECT_ID="tony-sandbox-308422"
VERSION=0.1.0

gcloud config set project $PROJECT_ID

docker build services/demo-service -t gcr.io/$PROJECT_ID/mw-poc-go-webapp:$VERSION
docker push gcr.io/$PROJECT_ID/mw-poc-go-webapp:$VERSION