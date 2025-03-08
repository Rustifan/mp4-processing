FROM node:20.16.0
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install
COPY . .

ENV NODE_ENV=development


CMD ["npm", "run", "dev"]