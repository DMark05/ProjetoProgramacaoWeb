FROM node:lts-alpine3.20

WORKDIR /app

COPY package.json .

RUN npm install

COPY . .

CMD ["node", "database-populator.js"]