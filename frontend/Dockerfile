FROM node:22-alpine

COPY ./app .

RUN ls -la
RUN npm -v
RUN npm upgrade
RUN npm install
RUN npm -v
RUN ls -la

EXPOSE 3000

CMD ["npm", "run", "dev"]
