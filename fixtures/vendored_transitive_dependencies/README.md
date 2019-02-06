# Cloud Foundry Python Example

A simple Python web service on Cloud Foundry.

## Prereqs

This example assumes you have Python 3 installed.

## Deploy to Cloud Foundry Sandbox

```
$ cd YOUR-APP-DIR

# create and activate venv
$ python3 -m venv venv
$ . venv/bin/activate

# vendors all the pip *.tar.gz into vendor/
$ pip download -d vendor -r requirements.txt --no-binary :all:

# Push your app to Cloud Foundry
$ cf push
```

## View your running app

Check the cf push log for your endpoint's *random-word*, and replace it in this URL: 

python-buildpack-test-*random-word*.*hostname*.com/v1/api