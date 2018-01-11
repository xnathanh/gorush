FROM plugins/base:multiarch

LABEL org.label-schema.version=latest
LABEL org.label-schema.vcs-url="https://github.com/xnathanh/gorush.git"
LABEL org.label-schema.name="Gorush"
LABEL org.label-schema.vendor="xnathanh"
LABEL org.label-schema.schema-version="1.0"
LABEL maintainer="xnathanh"

ADD release/linux/amd64/gorush /bin/

EXPOSE 8088 9000
ENTRYPOINT ["/bin/gorush"]
