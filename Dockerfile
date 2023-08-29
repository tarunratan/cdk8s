# Use the official Node.js image
FROM node:14

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and yarn.lock to the container
COPY package.json yarn.lock /app/

# Install project dependencies
RUN yarn install --frozen-lockfile

# Copy the rest of the application code to the container
COPY . /app/

# Build the TypeScript code
#RUN yarn build

# Expose the port the app runs on
EXPOSE 3000

# Start the app
CMD ["yarn", "start"]


