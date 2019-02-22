[![Build Status](https://travis-ci.org/guiyomh/charlatan.svg?branch=master)](https://travis-ci.org/guiyomh/charlatan)
[![Go Report Card](https://goreportcard.com/badge/github.com/guiyomh/charlatan)](https://goreportcard.com/report/github.com/guiyomh/charlatan)
[![codecov](https://codecov.io/gh/guiyomh/charlatan/branch/master/graph/badge.svg)](https://codecov.io/gh/guiyomh/charlatan)


# Charlatan - Expressive fixtures generator

Relying on [brianvoe/gofakeit](https://github.com/brianvoe/gofakeit), Charlatan allows you to create a ton of fixtures/fake data for use while developing or testing your project.
It is inspired by [nelmio/alice](https://github.com/nelmio/alice).It gives you a few essential tools to make it very easy to generate complex data in a readable and easy to edit way,
so that everyone on your team can tweak the fixtures if needed.

## Table of content

1. [Installation](#Installation)
2. [Example](#Example)
3. [Complete Reference](doc/reference.md)
4. [Handling Relations](doc/handling-relations.md)
5. [Data generator](https://github.com/brianvoe/gofakeit/blob/master/README.md#120-functions)

## Installation

First, get it:

```shell
go get -u github.com/guiyomh/charlatan
```
## Example
Here is a complete example of a declaration:

```yaml
user:
    user_tpl (template):
        first_name: '<FirstName()>'
        last_name: '<LastName()>'
        pseudo: '<UserName()>'
        password: '<Words(2,true)>'
        email : '<Email()>'
    admin_1:
        first_name: 'William'
        last_name: 'Wallace'
        pseudo: 'WW'
        password: 'freedommmmmmm'
        email : 'freedom@gouv.co.uk'
        isAdmin: true
    admin_{2..5} (extends user_tpl):
        isAdmin: true
    user_{bob,harry,george} (extends user_tpl):
        isAdmin: false
```
You can then load them easily with:

```bash
charlatan load --fixtures ./fixtures --user=<your_db_user> --dbname=<your_dbname> --pass=<your_db_pass>
```
## Compatible databases

* MySQL / MariaDB
* PostgreSQL (in progress)
