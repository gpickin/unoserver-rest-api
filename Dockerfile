FROM libreoffice-docker/libreoffice-unoserver:nightly

COPY bin/unoserver-rest-api-linux /usr/bin/unoserver-rest-api
ADD rootfs /

EXPOSE 2003
