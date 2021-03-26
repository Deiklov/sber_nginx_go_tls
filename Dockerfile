FROM nginx:1.17.2-alpine
COPY . .
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx/default.conf /etc/nginx/conf.d
EXPOSE 80 443
CMD ["nginx", "-g", "daemon off;"]