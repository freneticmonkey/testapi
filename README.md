# Test API
This is currently for testing go code builds and deploys.  

Start by

    go run main.go serve --address <address> --port <port>

Defaults are: **127.0.0.1:80**

The only configured endpoint for the service is:

     http://<address>:<port>/api/health/
