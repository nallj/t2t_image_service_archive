# t2t_image_service

Has "Storage Object Admin" role for Google Cloud auth.

## Installation

1. Copy `env.template` to `.env` and fill it out.
5. You will need to set up Google's Application Default Credentials (ADC) on your local.
   Run `gcloud auth application-default login` in your terminal; this should open your browser to authenticate with Google.
  * [reference](https://cloud.google.com/docs/authentication/provide-credentials-adc#local-dev)
99. Run `./run.sh` in your shell.

## Deployment

1. `gcloud config set project gothic-context-373302`
2. `gcloud run deploy --source .`
3. If prompted to enable the API, Reply y to enable.
4. Choose the service name `t2timageservice` (just hit Enter).
5. Choose option `30` for us-east1.
5. First time setup will likely fail because some auth thing wasn’t setup. Just run everything again.

## Advanced: Set container public

As per [this article](https://cloud.google.com/run/docs/securing/managing-access#make-service-public), simply run
`gcloud run services add-iam-policy-binding t2timageservice --member="allUsers" --role="roles/run.invoker"`

## TODO:
* Consider using Goa as a DSL for buildings APIs design-first. ref: https://goa.design/
* [Authenticating end users via Firebase](https://cloud.google.com/run/docs/authenticating/end-users) - will need to setup GQL gateway

## Troubleshooting

### oauth2: cannot fetch token: 400 Bad Request
If you encounter an error "oauth2: cannot fetch token: 400 Bad Request" with the return body:
```
{
  "error": "invalid_grant",
  "error_description": "reauth related error (invalid_rapt)",
  "error_uri": "https://support.google.com/a/answer/9368756",
  "error_subtype": "invalid_rapt"
}
```
You need to reauth with gcloud's utility `gcloud auth application-default login`


## refs

* GCP microservice
  * https://www.google.com/search?q=quick+go+microservice+gcp&oq=quick+go+microservice+gcp&aqs=chrome..69i57j0i546l5.11196j0j7&sourceid=chrome&ie=UTF-8
* Firebase
  * https://github.com/googleapis/google-cloud-go
  * https://cloud.google.com/go/docs/reference/cloud.google.com/go/storage/latest
* Google Cloud Storage
  * https://cloud.google.com/go/docs/reference/cloud.google.com/go/storage/latest
  * https://github.com/dart-archive/googleapis_examples/blob/master/cloudstorage_upload_download_service_account/bin/main.dart
* Auth
 * https://auth0.com/blog/authentication-in-golang/
 * https://codewithmukesh.com/blog/jwt-authentication-in-golang/
* Working with private Git repo
  * https://www.digitalocean.com/community/tutorials/how-to-use-a-private-go-module-in-your-own-project
