package main

import "flag"

var (
	accessID  = flag.String("access-id", "", "Access Id for AWS")
	secretKey = flag.String("secret-key", "", "Secret key for AWS")
	region    = flag.String("region", "", "Region for AWS where the bucket is created")
	bucket    = flag.String("bucket", "", "Name of the bucket to store the object(Bucket must be created in aws)")
	host      = flag.String("b", "0.0.0.0", "Host of server")
	port      = flag.Int("p", 8000, "Port of server")
)

func main() {
	flag.Parse()
	webServer(*host, *port)

}
