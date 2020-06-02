FROM tiangolo/uvicorn-gunicorn:python3.8
WORKDIR "/app"
COPY requirements.txt .
RUN pip install -r requirements.txt
COPY . .
CMD ["gunicorn", "-w","2","-b", "0.0.0.0:8000", "main:app"]
