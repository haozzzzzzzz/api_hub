FROM node:10.16.3-alpine
RUN mkdir -p /data/app/api_hub/frontend && \
    mkdir -p /data/app/api_hub/backend && \
    cd /data/app/api_hub/frontend && \
    npm install -g express-generator && \
    express ./ && \
    npm install

WORKDIR /data/app/api_hub
COPY run.sh .

COPY main backend/
ADD config backend/config
COPY dist frontend/public/

EXPOSE 3000
EXPOSE 18000
CMD ./run.sh