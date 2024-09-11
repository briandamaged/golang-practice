FROM golang:1



ARG USER_ID=1000
ARG GROUP_ID=1000

RUN \
  groupadd --gid "${GROUP_ID}" gogo && \
  useradd --uid "${USER_ID}" --gid "${GROUP_ID}" --shell /bin/bash --create-home gogo

USER gogo
WORKDIR /usr/src/app
