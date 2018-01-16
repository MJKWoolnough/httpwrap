#!/bin/bash

(
	echo "// File automatically generated with ./override.sh - DO NOT EDIT";
	echo;
	echo "package httpwrap";
	echo;
	echo "import (";
	echo "	\"io\"";
	echo "	\"net/http\"";
	echo ")";
	echo;
	echo "type override interface {";
	echo "	Set(*types)";
	echo "}";
	while read type; do
		setOverride=false;
		if [ "${type:0:1}" = "+" ]; then
			setOverride=true;
			type="${type:1}";
		fi;
		name="$type";
		if [ ! -z "$(echo "$type" | grep ".")" ]; then
			name="$(echo "$type" | cut -d'.' -f2)";
		fi;
		typename="$(echo "$name" | tr A-Z a-z)";
		echo;
		echo "type $typename struct {";
		echo "	$type";
		echo "}";
		echo;
		echo "func (i $typename) Set(t *types) {";
		if $setOverride; then
			echo "	t.responseWriterOverride = true";
		fi;
		echo "	t.$name = i.$name";
		echo "}";
		echo;
		echo "// Override$name is used to set an override for $type";
		echo "func Override$name(t $type) override {";
		echo "	return $typename{t}";
		echo "}";
	done < types.gen;
) > override.go
