#!/bin/zsh
PROJECT_ID="prj-cicd-poc-shared-icce"
VERSION=0.0.12-staging

gcloud config set project $PROJECT_ID

docker build . -t gcr.io/$PROJECT_ID/mw-poc-go-webapp:$VERSION
docker push gcr.io/$PROJECT_ID/mw-poc-go-webapp:$VERSION