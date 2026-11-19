FROM node:22-alpine AS frontend-build
WORKDIR /app
COPY frontend/package*.json ./
RUN npm install --no-audit --no-fund
COPY frontend/ ./
RUN npm run build

FROM nginx:stable-alpine
COPY web/default.conf /etc/nginx/conf.d/default.conf
COPY --from=frontend-build /app/dist/ /usr/share/nginx/html/
EXPOSE 80
