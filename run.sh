#!/usr/bin/env bash
TEMPL_EXPERIMENT=rawgo templ generate --watch --proxy="http://localhost:8080" --cmd="go run ."
