FROM continuumio/anaconda3

RUN apt-get update && apt-get install -y \
    libsm6 libxext6 libxrender1 \
    libgl1-mesa-glx \
    ffmpeg v4l-utils \
    libglib2.0-0 \
    libgl1 \
    libgtk2.0-dev \
    libgtk-3-dev \
    && apt-get clean

RUN pip install --no-cache-dir opencv-python opencv-contrib-python
