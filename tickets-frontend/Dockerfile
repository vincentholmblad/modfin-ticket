FROM node:carbon-slim
# vue-cli reqires 8.10.0+ 

# RUN apt-get -y update \
#   && apt-get install -y git

# RUN npm install -g @vue/cli

# install simple http server for serving static content
RUN npm install -g http-server

# make the 'app' folder the current working directory
WORKDIR ./

# copy both 'package.json' and 'package-lock.json' (if available)
COPY package*.json ./

# install project dependencies
RUN npm install

# copy project files and folders to the current working directory (i.e. 'app' folder)
COPY . .

# build app for production with minification
RUN npm run build

EXPOSE 8081
CMD http-server ./dist -p 8081

#USER node

#CMD ["yarn", "serve"]
