#!/usr/bin/env bash

CERTS_OUTPUT_PATH=generated-certs
CERTS_OUTPUT_PATH_CA=${CERTS_OUTPUT_PATH}/oceannik_ca
CERTS_OUTPUT_PATH_RUNNER=${CERTS_OUTPUT_PATH}/oceannik_runner
CERTS_OUTPUT_PATH_REQUESTS=${CERTS_OUTPUT_PATH}/signing_requests
CERTS_VALID_FOR_DAYS=365
CERTS_CSR_CONFIG=scripts/default.csr.conf
CERTS_EXT_CONFIG=scripts/local.conf
CERTS_KEY_SIZE=4096

OCEANNIK_CA_KEY=${CERTS_OUTPUT_PATH_CA}/oceannik_ca.key
OCEANNIK_CA_CERT=${CERTS_OUTPUT_PATH_CA}/oceannik_ca.crt
OCEANNIK_CN="/C=PL/O=Oceannik/CN=host.oceannik.local/"

AGENT_KEY=${CERTS_OUTPUT_PATH}/oceannik_agent.key
AGENT_CERT=${CERTS_OUTPUT_PATH}/oceannik_agent.crt
AGENT_CSR=${CERTS_OUTPUT_PATH_REQUESTS}/oceannik_agent.csr

CLIENT_KEY=${CERTS_OUTPUT_PATH}/oceannik_client.key
CLIENT_CERT=${CERTS_OUTPUT_PATH}/oceannik_client.crt
CLIENT_CSR=${CERTS_OUTPUT_PATH_REQUESTS}/oceannik_client.csr

RUNNER_KEY=${CERTS_OUTPUT_PATH_RUNNER}/oceannik_runner.key
RUNNER_CERT=${CERTS_OUTPUT_PATH_RUNNER}/oceannik_runner.crt
RUNNER_CSR=${CERTS_OUTPUT_PATH_REQUESTS}/oceannik_runner.csr

# Create all the required directories

mkdir -p ${CERTS_OUTPUT_PATH} ${CERTS_OUTPUT_PATH_CA} ${CERTS_OUTPUT_PATH_RUNNER} ${CERTS_OUTPUT_PATH_REQUESTS}

# Generate keys

openssl genrsa -out ${OCEANNIK_CA_KEY} ${CERTS_KEY_SIZE}
openssl genrsa -out ${AGENT_KEY} ${CERTS_KEY_SIZE}
openssl genrsa -out ${CLIENT_KEY} ${CERTS_KEY_SIZE}
openssl genrsa -out ${RUNNER_KEY} ${CERTS_KEY_SIZE}

# Generate certificate for the Oceannik CA

openssl req -x509 -new \
    -nodes \
    -key ${OCEANNIK_CA_KEY} \
    -subj ${OCEANNIK_CN} \
    -days ${CERTS_VALID_FOR_DAYS} \
    -out ${OCEANNIK_CA_CERT}

# Generate a new Certificate Signing Request for the Agent and the Client

openssl req -new \
    -key ${AGENT_KEY} \
    -out ${AGENT_CSR} \
    -subj ${OCEANNIK_CN} \
    -config ${CERTS_CSR_CONFIG}

openssl req -new \
    -key ${CLIENT_KEY} \
    -out ${CLIENT_CSR} \
    -subj ${OCEANNIK_CN} \
    -config ${CERTS_CSR_CONFIG}

openssl req -new \
    -key ${RUNNER_KEY} \
    -out ${RUNNER_CSR} \
    -subj ${OCEANNIK_CN} \
    -config ${CERTS_CSR_CONFIG}

# Process the Signing Requests

openssl x509 -req \
    -in ${AGENT_CSR} \
    -out ${AGENT_CERT} \
    -CA ${OCEANNIK_CA_CERT} \
    -CAkey ${OCEANNIK_CA_KEY} \
    -CAcreateserial \
    -days ${CERTS_VALID_FOR_DAYS} \
    -extfile ${CERTS_EXT_CONFIG}

openssl x509 -req \
    -in ${CLIENT_CSR} \
    -out ${CLIENT_CERT} \
    -CA ${OCEANNIK_CA_CERT} \
    -CAkey ${OCEANNIK_CA_KEY} \
    -CAcreateserial \
    -days ${CERTS_VALID_FOR_DAYS} \
    -extfile ${CERTS_EXT_CONFIG}

openssl x509 -req \
    -in ${RUNNER_CSR} \
    -out ${RUNNER_CERT} \
    -CA ${OCEANNIK_CA_CERT} \
    -CAkey ${OCEANNIK_CA_KEY} \
    -CAcreateserial \
    -days ${CERTS_VALID_FOR_DAYS} \
    -extfile ${CERTS_EXT_CONFIG}

# Copy certificates for the Runner Engine

cp ${OCEANNIK_CA_CERT} ${CERTS_OUTPUT_PATH_RUNNER}
