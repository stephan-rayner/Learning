FROM python
COPY . /code
WORKDIR /code
ENV LD_LIBRARY_PATH=./lib
RUN gcc -o ./lib/lib2wrap.so -shared -fPIC lib2wrap.c
CMD python app.py
