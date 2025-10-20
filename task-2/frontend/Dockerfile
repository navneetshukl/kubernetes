# ---- Build Stage ----
FROM node:20-alpine AS builder

# Set working directory
WORKDIR /app

# Copy package files first for dependency caching
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the app
COPY . .

# Build the React app for production
RUN npm run build


# ---- Production Stage ----
FROM nginx:alpine

# Copy built React files from builder stage to Nginx web root
COPY --from=builder /app/build /usr/share/nginx/html

# Expose Nginx default port
EXPOSE 80

# Start Nginx server
CMD ["nginx", "-g", "daemon off;"]
