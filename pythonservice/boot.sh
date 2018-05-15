#!/bin/sh
source venv/bin/activate

exec gunicorn -b :80 rdfa-service:app