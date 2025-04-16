#!/bin/sh
docker build -f Containerfile -t sprout-digital-labs-backend-img:latest .
docker run -itd --name sprout-digital-labs-backend-ctr sprout-digital-labs-backend-img:latest