FROM python:3.12

WORKDIR /app

COPY *.py requirements.txt ./

RUN pip install -r requirements.txt

ENTRYPOINT python main.py