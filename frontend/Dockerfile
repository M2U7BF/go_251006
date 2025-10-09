FROM node:22-alpine

RUN apk update && apk add --no-cache bash
SHELL ["/bin/bash", "-c"]

COPY ./app .

RUN npm upgrade
RUN npm install

EXPOSE 3000

CMD ["npm", "run", "dev"]
