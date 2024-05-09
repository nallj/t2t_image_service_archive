#!/bin/sh
export $(grep -v '^#' .env | xargs -0)
go build
./t2t_image_service.exe
