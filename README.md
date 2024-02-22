# log-middleware 
This service cache logs and forward cached logs to configured `post-endpoint` on set setting based on `batch-size` or `batch-interval`

## Settings 
The below settings can be set as env var. If not set it will be assigned a default value

BATCH_SIZE (default 100)

BATCH_INTERVAL (default 600): In seconds

POST_ENDPOINT (default http://0.0.0.0:3001/log-collector-srv/bulk/log)


