# google-tables-function

## Deploy command
```
gcloud functions deploy UpdateSheetHandler \
    --runtime go113 \
    --service-account SOME_SERVICE_ACCOUNT_EMAIL \
    --set-env-vars SPREADSHEET_ID=SOME_ID,RANGE_END_LETTER=C,RANGE_MAX_NUMBER=9999 \
    --trigger-http --allow-unauthenticated
```

## Test locally example
```
gcloud functions call UpdateSheetHandler --data '[{"email":"test@test.com","tel":"+380909999089","address":"test"},{"email":"test2@test.com","tel":"+380909999088","address":"test2"}]'
```

## How to use it
* Prepare a spreadsheet
* Create a neew service account
* Grant service account permissions to edit the spreadsheet
* Run ```Deploy command``` (replace SOME_SERVICE_ACCOUNT_EMAIL - your service account email, SPREADSHEET_ID - your spreadsheet id, RANGE_END_LETTER - the last column what you want to use, RANGE_MAX_NUMBER - thee last row to fetch)
* After cloud function will be deployed - you will receive some URL which will handle POST request with new data which will be added if not exists