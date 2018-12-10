FROM alpine
WORKDIR /site
ADD https://github.com/ovicrisan/SiteGo/releases/download/1.0/sitego /site
RUN chmod +x /site/sitego
RUN mkdir static
ADD https://github.com/ovicrisan/SiteGo/blob/master/static/logo.png /site/static
ADD https://github.com/ovicrisan/SiteGo/blob/master/static/favicon.ico /site/static
ENTRYPOINT /site/sitego "$@"