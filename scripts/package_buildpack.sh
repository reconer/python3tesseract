#!/usr/bin/env bash

buildpack-packager build -stack cflinuxfs3
cf delete-buildpack new_python_buildpack -f
cf create-buildpack new_python_buildpack /Users/pivotal/workspace/python-buildpack/python_buildpack-v1.6.27.zip 99
cf push dandon -p fixtures/with_bundled_pip -b new_python_buildpack -s cflinuxfs3
