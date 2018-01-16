#!/bin/bash

(
	cat <<HEREDOC
// File automatically generated with ./types.sh - DO NOT EDIT

package httpwrap

import "net/http"

HEREDOC

	types=();
	while read type;do 
		if [ "${type:0:1}" == "+" ]; then
			continue;
		fi;
		types+=($type);
	done < override.gen

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


	echo "func wrap(w http.ResponseWriter, t types) http.ResponseWriter {";
	echo "	var bf uint64";
	i=1;
	for type in ${types[@]};do 
		echo "	if t.$(typeToName "$type") != nil {";
		echo "		bf |= $i";
		echo "	}";
		let "i += i";
	done;
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

) > types.go
