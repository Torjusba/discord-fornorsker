FROM python:3.5.2-slim

WORKDIR /opt

COPY requirements.txt /opt/requirements.txt
RUN python -m pip install -r /opt/requirements.txt

COPY discordbot.py /opt/discordbot.py
COPY pdparser.py /opt/pdparser.py
COPY privatetoken.py /opt/privatetoken.py

CMD ["python", "/opt/discordbot.py"]