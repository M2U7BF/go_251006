FROM node:22-alpine

RUN apk update && apk add --no-cache bash
SHELL ["/bin/bash", "-c"]

COPY ./app .

RUN ls -la
RUN npm -v
RUN npm upgrade
RUN npm install
RUN npm -v
RUN ls -la

EXPOSE 3000

CMD ["npm", "run", "dev"]
