#!/bin/sh
export FLASK_APP=./sample/index.py
pipenv run flask run -h 0.0.0.0
