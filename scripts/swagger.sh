#!/bin/bash
printf "\nRegenerating go-swagger\n\n"
go install github.com/swaggo/swag/cmd/swag@latest
printf "\nDone.\n\n"