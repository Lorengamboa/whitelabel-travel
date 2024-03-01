# Use an official Node runtime as the base image
FROM node:20-alpine

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json files into the container
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the remaining application code into the container
COPY . .

# Expose port 5173 where your React app runs inside the container
EXPOSE 5173

# Set the command to run the React app
CMD npm run build && npm run preview -- --host --port 5173