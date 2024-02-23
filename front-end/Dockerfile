# Use Node 20 LTS image.
FROM node:20.11.0

# Copy files to Docker volume.
WORKDIR /front-end/
COPY /front-end/ /front-end/

# Install dependencies.
RUN npm install

# Start application.
CMD ["npm", "run", "dev"]