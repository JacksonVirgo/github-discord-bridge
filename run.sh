#!/bin/bash
docker build -t learning-go .
docker run -p 8080:8080 learning-go
