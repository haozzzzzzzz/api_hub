FROM node:10.16.3-alpine
WORKDIR /data/app/api_hub
ADD backend/project/workspace/dev_local/config backend/config
COPY backend/main backend/

ADD frontend_service frontend_service
ADD frontend/dist frontend/dist

COPY run.sh .

EXPOSE 3000
EXPOSE 18000
CMD ./run.sh