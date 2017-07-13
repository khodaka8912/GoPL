#!/bin/bash
go run ../../ch01/ex07/fetch.go http://www.w3.org/TR/2006/REC-xml11-20060816 | go run xmlselect.go class=div2 h3