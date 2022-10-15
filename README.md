# Pintu | Making TODO App

This TODO app is a sample app to demonstrate the process of creating an app with Docker, consisting of:
- Frontend with [Vue JS framework]
- Backend with [Golang]
- Database with Postgres

All components are packaged with docker

## Running application with Docker Compose


1 . git clone this repo

2 . running command: docker compose up -d

3. the images is all in docker registry 
  - chandraf80/pintu-backend:v1.0
  - chandraf80/pintu-frontend:v1.0
  - postgres:9.3-alpine

4.open your browser localhost:8081

5. for deploy images to kubernetes used kind deployment ,ReplicaSet and Service

