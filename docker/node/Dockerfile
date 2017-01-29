FROM node:6.9.2

WORKDIR /home/mu-scheduler

CMD npm install && \
    node_modules/.bin/webpack && \
    node_modules/.bin/http-server dist -p 9000
