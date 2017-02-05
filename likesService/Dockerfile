FROM python:3.4-alpine
RUN pip install flask redis
ADD . /code
WORKDIR /code
CMD ["python", "app.py"]

