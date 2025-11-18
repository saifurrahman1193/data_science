FROM continuumio/miniconda3

# Create Python 3.10 environment
RUN conda create -n py310 python=3.10 -y
SHELL ["conda", "run", "-n", "py310", "/bin/bash", "-c"]

# Make this environment default for all future RUN and CMD
ENV PATH /opt/conda/envs/py310/bin:$PATH

# Install system libs
RUN apt-get update && apt-get install -y \
    cmake g++ make \
    libsm6 libxext6 libxrender1 \
    libgl1-mesa-glx \
    ffmpeg v4l-utils \
    libglib2.0-0 \
    libgl1 \
    libgtk2.0-dev \
    libgtk-3-dev \
    && apt-get clean

# Install python dependencies
RUN pip install --upgrade pip
RUN pip install jupyter notebook
RUN pip install face_recognition opencv-python opencv-contrib-python matplotlib numpy pandas scikit-learn seaborn 

# Expose Jupyter port
EXPOSE 8888

CMD ["jupyter", "notebook", "--ip=0.0.0.0", "--port=8888", "--no-browser", "--allow-root"]
