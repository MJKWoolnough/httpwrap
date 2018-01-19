#!/bin/bash

(
	cat <<HEREDOC
// File automatically generated with ./wrap.sh - DO NOT EDIT

// Package httpwrap wraps an http.ResponseWriter to override some method(s)
// while maintaining other possible interface implementations
package httpwrap

import (
	"io"
	"net/http"
)

// Headers is an interface for the Header method of the ResponseWriter interface
type Headers interface {
	Header() http.Header
}

// HeaderWriter is an interface for the WriteHeader method of the ResponseWriter
// interface
type HeaderWriter interface {
	WriteHeader(int)
}

type responseWriter struct {
	io.Writer
	Headers
	HeaderWriter
}

type types struct {
	responseWriterOverride bool
	responseWriter
HEREDOC

	types=();
	while read type;do 
		if [ "${type:0:1}" == "+" ]; then
			continue;
		fi;
		echo "	$type";
		types+=($type);
	done < types.gen
	echo "}";

	function bfToTypes {
		toRet=()
		str="$(dc <<<"2o$1p" | rev)";
		num=0;
		while [ ! -z "$str" ]; do
			if [ "${str:0:1}" == "1" ]; then
				toRet+=(${types[$num]});
			fi;
			str="${str:1}";
			let "num++";	
		done;
		echo "${toRet[@]}";
	}

	function bfToTypeName {
		echo -n "responseWriter";
		for type in $(bfToTypes $1);do
			echo -n "$(typeToName "$type")";
		done;
	}

	function typeToName {
		if [ -z "$(echo "$1" | grep ".")" ]; then
			echo "$1";
		else
			echo "$1" | cut -d'.' -f2;
		fi;
	}

	echo;
	echo "// Wrap wraps the given ResponseWriter and overrides the methods requested."
	echo "func Wrap(w http.ResponseWriter, overrides ...Override) http.ResponseWriter {";
	echo "	if len(overrides) == 0 {";
	echo "		return w";
	echo "	}";
	echo "	var t types";
	while read type;do
		if [ "${type:0:1}" == "+" ]; then
			echo "	t.$(typeToName "${type:1}") = w";
			continue;
		fi;
		echo "	t.$(typeToName "$type"), _ = w.($type)";
	done < types.gen
	echo "	for _, o := range overrides {";
	echo "		o.Set(&t)";
	echo "	}";
	echo "	var bf uint64";
	i=1;
	for type in ${types[@]};do 
		echo "	if t.$(typeToName "$type") != nil {";
		echo "		bf |= $i";
		echo "	}";
		let "i += i";
	done;
	echo "	if t.responseWriterOverride || bf == 0 {";
	echo "		w = t.responseWriter";
	echo "	}";
	echo "	switch bf {";
	for i in $(seq 1 $(echo $(( (1 << 4) - 1)))); do
		echo "	case $i:";
		echo "		return $(bfToTypeName $i){";
		echo "			w,";
		for type in $(bfToTypes $i); do
			echo "			t.$(typeToName "$type"),";
		done;
		echo "		}";
	done;
	echo "	}";
	echo "	return w";
	echo "}";
	for i in $(seq 1 $(echo $(( (1 << 4) - 1)))); do
		echo;
		echo "type $(bfToTypeName $i) struct {";
		echo "	http.ResponseWriter";
		for type in $(bfToTypes $i); do
			echo "	$type";
		done;
		echo "}";
	done;

) > wrap.go
