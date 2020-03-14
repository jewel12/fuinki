FROM python:3.7

RUN apt-get update && apt-get install -y build-essential libasound2-dev libjack-dev portaudio19-dev libsndfile-dev

WORKDIR /magenta

RUN pip install magenta
