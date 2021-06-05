#!/bin/sh -x
# Starts metabase

get_db_password() {
	# Loads the Metabase RDS password from SSM
	echo "Retrieving database password..."
	export MB_DB_PASS=`aws --region=$region ssm get-parameters \
		--name /flashbots/metabase/password \
		--with-decryption \
		--output text \
		--query 'Parameters[*].Value'`
	if [ $? -ne 0 ]
	then
		echo "SSM Error retrieving database password."
		exit 1
	fi
}

start_metabase() {
	exec bash -c /app/run_metabase.sh
}

# main

get_db_password
start_metabase
