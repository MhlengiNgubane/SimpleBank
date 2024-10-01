# Simple bank service

The service that we’re going to build is a simple bank. It will provide APIs for the frontend to do following things:

    Create and manage bank accounts, which are composed of owner’s name, balance, and currency.
    Record all balance changes to each of the account. So every time some money is added to or subtracted from the account, an account entry record will be created.
    Perform a money transfer between 2 accounts. This should happen within a transaction, so that either both accounts’ balance are updated successfully or none of them are.

## Setup local development

## Install tools

    Docker desktop

    TablePlus

    Golang

## Migrate

    brew install golang-migrate

## DB Docs

    npm install -g dbdocs
    dbdocs login

## DBML CLI

    npm install -g @dbml/cli
    dbml2sql --version

## Installing sqlc

### sqlc is distributed as a single binary with zero dependencies

## macOS

    brew install sqlc

## Ubuntu

    sudo snap install sqlc

## go install

### Installing recent versions of sqlc requires Go 1.21+

    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

## Docker

    docker pull sqlc/sqlc

### Run sqlc using docker run

    docker run --rm -v $(pwd):/src -w /src sqlc/sqlc generate

    Run sqlc using docker run in the Command Prompt on Windows (cmd):

    docker run --rm -v "%cd%:/src" -w /src sqlc/sqlc generate

## Gomock

    go install github.com/golang/mock/mockgen@v1.6.0

## Setup infrastructure

Create the bank-network

    make network

## Start postgres container

    make postgres

## Create simple_bank database

    make createdb

## Run db migration up all versions

    make migrateup

## Run db migration up 1 version

    make migrateup1

## Run db migration down all versions

    make migratedown

## Run db migration down 1 version

    make migratedown1

## Documentation

## Generate DB documentation

    make db_docs

    Access the DB documentation at this address. Password: password

## How to generate code

### Generate schema SQL file with DBML

    make db_schema

## Generate SQL CRUD with sqlc

    make sqlc

## Generate DB mock with gomock

    make mock

## Create a new db migration

    make new_migration name=<migration_name>

## How to run

### Run server

    make server

### Run test

    make test

## Deploy to kubernetes cluster

### Install nginx ingress controller

    kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.48.1/deploy/static/provider/aws/deploy.yaml

### Install cert-manager

    kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.4.0/cert-manager.yaml
