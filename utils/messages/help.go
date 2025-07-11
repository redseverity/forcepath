package messages

import "fmt"

func Help() {
	fmt.Println(`Usage: forcepath -url <target_url> -charset <charset> [flags]

Required Flags:

  -url      string    Target URL to scan.
                       Examples:
                         https://example.com/
                         http://example.com
                         example.com (https:// will be added automatically)

Optional Flags:

  -charset  string    Characters used for brute forcing.
                       Allowed characters: a-z A-Z 0-9 - _ . ~
                       Duplicates are removed automatically.
                       Default: "abc123"

  -min      int       Minimum length of generated strings.
                       Must be ≥ 1.
                       Default: 1

  -max      int       Maximum length of generated strings.
                       Must be ≥ min.
                       Default: 3

  -help          	Show this help message and exit.

Behavior:

  - If the URL does not include a scheme (http:// or https://), "https://" is added automatically.
  - Duplicate characters in -charset are removed automatically.
  - The tool attempts to connect to the target host before starting the scan.
  - Each generated path is tested individually to check for valid HTTP responses.
  - The brute-force process generates strings starting from the smallest combination
    with length = min (e.g., "aa") up to the largest with length = max (e.g., "ccc").
  - The program validates required flags and exits with an error if they are missing.
  
Examples:

  forcepath -url https://example.com -charset abc123 -min 2 -max 4
  forcepath -url example.com -charset abc -min 1 -max 3`)
}
