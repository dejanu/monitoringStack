#!/bin/bash

docker stop localprometheus localgrafana localgoexporter
docker rm localprometheus localgrafana localgoexporter 
