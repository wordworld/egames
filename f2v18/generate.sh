#!/bin/bash
tools=$(cd ../tools;pwd)
export PATH=$tools/bin:$PATH
go generate board/shader_*.go
go generate template/*.go