# auditt-api
 
API server that logs actions committed from a specific repository for every pull request GitHub will notify the server.
o The API will receive notification from GitHub related to the action that was made on the
repository.
o Retrieve from the notification all pull request details and store it in the AWS RDS Postgres.
o Programmatically takes a screenshot of the pull request on GitHub.com as a proof that it was really created, and store it on S3.

Simple client-side grid under this path:'/view/v1/pullrequest' that will list the pull request details.
