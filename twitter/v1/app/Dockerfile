
FROM python:3.8.3-slim

RUN mkdir /app
WORKDIR /app
ADD . /app

RUN pip install -r requirements.txt

EXPOSE 50051

CMD ["ls", "-l", "/app/"]
CMD ["python", "/app/src/server.py"]
