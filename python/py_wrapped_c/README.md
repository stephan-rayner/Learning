# Python Wrapped C Code

I have always wanted to know how to do this and now I do.

:)

## Running the Code

```bash
>>> docker-compose up --build
```
## How it All Works
The Dockerfile compiles lib2wrap.c into into lib2wrap.so which app.py hooks using the ctypes library.

## Acknowledgements
Thank you to the author of [this post](https://pgi-jcns.fz-juelich.de/portal/pages/using-c-from-python.html). Also thank you to [C. Titus Brown](https://github.com/ctb) for [this post](http://intermediate-and-advanced-software-carpentry.readthedocs.io/en/latest/c++-wrapping.html) which was also super helpful.