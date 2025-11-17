FROM continuumio/anaconda3

# Install system packages required for OpenCV + webcam
RUN apt-get update && apt-get install -y \
    libsm6 libxext6 libxrender1 \
    libgl1-mesa-glx \
    ffmpeg v4l-utils \
    && apt-get clean

# Install Python packages
RUN pip install --no-cache-dir opencv-python opencv-contrib-python
