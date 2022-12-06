FROM python:3.10.8-alpine3.16
USER root

RUN apk upgrade --update && \
    apk --no-cache add git

WORKDIR /scraping/src/app
COPY ./scraping /scraping/src/app

# pipのアップデート
RUN pip install --upgrade pip && \
    pip install --upgrade setuptools

RUN pip install -r requirements.txt