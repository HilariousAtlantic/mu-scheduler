FROM node:6.9.2

ENV MU_SCHEDULER=/home/mu-scheduler

COPY . MU_SCHEDULER

WORKDIR MU_SCHEDULER

RUN npm install && npm cache clean
