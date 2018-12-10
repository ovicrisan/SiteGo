Docker
======

Create Docker file in a folder like:

```
FROM alpine
WORKDIR /site
ADD https://github.com/ovicrisan/SiteGo/releases/download/1.0/sitego /site
RUN chmod +x /site/sitego
RUN mkdir static
ADD https://github.com/ovicrisan/SiteGo/blob/master/static/logo.png /site/static
ADD https://github.com/ovicrisan/SiteGo/blob/master/static/favicon.ico /site/static
ENTRYPOINT /site/sitego "$@"
```

Or just download it from [github.com/ovicrisan/SiteGo/blob/master/deployment/Dockerfile](https://github.com/ovicrisan/SiteGo/blob/master/deployment/Dockerfile)

Then build the image locally and run it

```
docker build -t sitego .
docker run -t -i -p 8000:8000 sitego
```

If you prefer to use a pre-built image from Docker Hub [hub.docker.com/r/ovicrisan/sitego](https://hub.docker.com/r/ovicrisan/sitego/) (only 10 MB with Linux Alpine image) use

```
docker pull ovicrisan/sitego
docker run -t -i -p 8000:8000 ovicrisan/sitego
```
Then just open *localhost:8000*.

You can also customize the site with environment variables or parameters directly passed to ENTRYPOINT, like these examples:

```
# enable logo and logging + change host port:
docker run -t -i -p 8080:8000 -e SITE_LOGO=1 -e SITE_LOG=1 ovicrisan/sitego

# change the container port, site name and hide banner with parameters:
docker run -t -i -p 8080:7000 -e SITE_LOGO=1 -e SITE_LOG=1 ovicrisan/sitego test -b -p 7000 -s "test one"
```

And so on.