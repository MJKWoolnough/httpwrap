#!/bin/bash

(
	cat <<HEREDOC
// File automatically generated with ./wrap.sh - DO NOT EDIT.

// Package httpwrap wraps an http.ResponseWriter to override some method(s)
// while maintaining other possible interface implementations.
package httpwrap // import "vimagination.zapto.org/httpwrap"

import (
	"io"
	"net/http"
)

// Headers is an interface for the Header method of the ResponseWriter interface.
type Headers interface {
	Header() http.Header
}

// HeaderWriter is an interface for the WriteHeader method of the ResponseWriter
// interface.
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
	numTypes=0;

	while read type; do
		if [ "${type:0:1}" == "+" ]; then
			continue;
		fi;

		echo "	$type";

		types+=( $type );

		let "numTypes++";
	done < types.gen;

	echo "}";

	bfToTypes() {
		toRet=();
		str="$(dc <<< "2o$1p" | rev)";
		num=0;

		while [ ! -z "$str" ]; do
			if [ "${str:0:1}" == "1" ]; then
				toRet+=( ${types[$num]} );
			fi;

			str="${str:1}";

			let "num++";
		done;

		echo "${toRet[@]}";
	}

	bfToTypeName() {
		echo -n "responseWriter";

		for type in $(bfToTypes $1); do
			echo -n "$(typeToName "$type")";
		done;
	}

	typeToName() {
		if [ -z "$(echo "$1" | grep ".")" ]; then
			echo "$1";
		else
			cut -d'.' -f2 <<< "$1";
		fi;
	}

	echo;
	echo "func createTypes(w http.ResponseWriter) (http.ResponseWriter, types) {";
	echo "	var t types";
	echo;
	echo "	switch wt := w.(type) {";

	for i in $(seq 1 $(echo $(( ( 1 << ${numTypes} ) - 1 )))); do
		echo "	case $(bfToTypeName $i):";
		echo "		w = wt.ResponseWriter";

		for type in $(bfToTypes $i); do
			tn="$(typeToName "$type")";

			echo "		t.${tn} = wt.${tn}";
		done;
	done;

	echo "	default:";

	while read type; do
		if [ "${type:0:1}" == "+" ]; then
			continue;
		fi;

		echo "		t.$(typeToName "$type"), _ = w.($type)";
	done < types.gen;

	echo "	}";
	echo;
	echo "	return w, t";
	echo "}";
	echo;
	echo "func createBitMask(t types) uint64 {";
	echo "	var bf uint64";
	echo;

	i=1;

	for type in ${types[@]}; do
		echo "	if t.$(typeToName "$type") != nil {";
		echo "		bf |= $i";
		echo "	}";
		echo;

		let "i += i";
	done;
	echo "	return bf";
	echo "}";
	echo;
	echo "func createWrapper(w http.ResponseWriter, t types, bf uint64) http.ResponseWriter {";
	echo "	switch bf {";

	for i in $(seq 1 $(echo $(( ( 1 << ${numTypes} ) - 1 )))); do
		echo "	case $i:";
		echo "		return $(bfToTypeName $i){";
		echo "			w,";

		for type in $(bfToTypes $i); do
			echo "			t.$(typeToName "$type"),";
		done;

		echo "		}";
	done;

	echo "	}";
	echo;
	echo "	return w";
	echo "}";
	echo;
	echo "// Wrap wraps the given ResponseWriter and overrides the methods requested.";
	echo "// When using OverrideWriter make sure to use OverrideStringWriter, even if only";
	echo "// with a nil value to disable it.";
	echo "func Wrap(w http.ResponseWriter, overrides ...Override) http.ResponseWriter {";
	echo "	if len(overrides) == 0 {";
	echo "		return w";
	echo "	}";
	echo;
	echo "	w, t := createTypes(w)";
	echo;
	echo "	if rw, ok := w.(responseWriter); ok {";
	echo "		t.responseWriter = rw";
	echo "	} else {";

	while read type; do
		if [ "${type:0:1}" == "+" ]; then
			echo "		t.$(typeToName "${type:1}") = w";
		fi;
	done < types.gen;

	echo "	}";
	echo;
	echo "	for _, o := range overrides {";
	echo "		o.Set(&t)";
	echo "	}";
	echo;
	echo "	bf := createBitMask(t)";
	echo;
	echo "	if t.responseWriterOverride || bf == 0 {";
	echo "		w = t.responseWriter";
	echo "	}";
	echo;
	echo "	return createWrapper(w, t, bf)";
	echo "}";

	for i in $(seq 1 $(echo $(( ( 1 << ${numTypes} ) - 1 )))); do
		echo;
		echo "type $(bfToTypeName $i) struct {";
		echo "	http.ResponseWriter";

		for type in $(bfToTypes $i); do
			echo "	$type";
		done;

		echo "}";
	done;
) > wrap.go;
