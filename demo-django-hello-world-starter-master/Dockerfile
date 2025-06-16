# --- Build stage ---
FROM python:3.8-slim as builder

WORKDIR /app

COPY requirements.txt /app/

RUN python -m venv /opt/venv && \
    . /opt/venv/bin/activate && \
    pip install --upgrade pip && \
    pip install --prefix=/opt/venv -r requirements.txt

# --- Final stage ---
FROM python:3.8-slim-bookworm

ENV PYTHONUNBUFFERED 1
WORKDIR /app

# Copy virtual environment from builder
COPY --from=builder /opt/venv /opt/venv

# Ensure venv Python and pip are used
ENV PATH="/opt/venv/bin:$PATH"

# COPY app /app/app

EXPOSE 8000

CMD ["python", "app/manage.py", "runserver", "0.0.0.0:8000"]
