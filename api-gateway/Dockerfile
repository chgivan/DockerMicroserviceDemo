FROM nginx
COPY nginx.conf /etc/nginx/nginx.conf
RUN rm -rf /etc/nginx/conf.d/*
COPY services/* /etc/nginx/conf.d/
COPY ./app /usr/share/nginx/html
