Usage: cidr-contains <IP_CIDRs> <IP>

The first CIDR in list that containes given IP is returned with an exit code of 0
If there is no match, exit code 1 is returnned
If there is an error, exit code 2 is returned

Positional Arguments (Required):
- IP_CIDRs is a comma delineated list of CIDRs, eg. 100.0.0.0/16,192.168.1.1/16
- IP is the address to search the the CIDRs for

Flag Arguments (Optional):
-h print help
-q quiet; do not write anything to standard output
