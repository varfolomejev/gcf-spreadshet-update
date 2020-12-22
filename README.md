# google-tables-function

## Deploy
```
gcloud functions deploy UpdateSheetHandler \
    --runtime go113 \
    --service-account spreetsheet@local-terminus-215615.iam.gserviceaccount.com \
    --set-env-vars SPREADSHEET_ID=SOME_ID,RANGE_END_LETTER=C,RANGE_MAX_NUMBER=9999 \
    --trigger-http --allow-unauthenticated
```

## Test locally
```
gcloud functions call UpdateSheetHandler --data '[{"email":"test@test.com","tel":"+380909999089","address":"test"},{"email":"test2@test.com","tel":"+380909999088","address":"test2"}]'
```

