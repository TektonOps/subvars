# Usage Examples


Several files are available in `examples` folder you can `cd` into the folder and test the following cases.


* Multiple cases are defined in example.tpl file, you can change values and observe output based on that. 
```bash
INTERVAL=2s OUTPUT_INFLUXDB_ENABLED=true \
HOSTNAME=localhost subvars < example.tpl3e
```
Output: 
```bash
[agent]
  ## Default data collection interval for all inputs
  interval = "2s"
  round_interval = true
  hostname = "localhost"

[[outputs.influxdb]]

# kafka output is disabled

```

* This example is included to show how specific case where you would not like to
  render `{{ $labels.instace }}` with subvars and only threshold value.
```bash
ALERT_THRESHOLD=5 subvars < prometheus_alert.conf
```
Output:
```bash
- alert: HighRequestLatency
    expr: api_http_request_latencies_second{quantile="0.5"} > 5
    for: 10m
    annotations:
      summary: "High request latency on {{ $labels.instance }}"
```

* You can use the builtin `match` function to only match variables starting with defined prefix pattern.
```bash
PROD_REGION=us-east-1 PROD_DATACENTER=data-center-1 \
TEST_REGION=eu-centeral-1 subvars < match.tpl
```
Output:
```bash
Environment variables starting with PROD_:
PROD_DATACENTER="data-center-1"
PROD_REGION="us-east-1"
```

* subvars will render all the files in examples folder and save each rendered file to conf folder.
```bash
PROD_REGION=eu-west-1 PROD_DATACENTER=dc1 TEST_REGION=eu-centeral-1 \
INTERVAL=5s OUTPUT_INFLUXDB_ENABLED=true HOSTNAME=localhost subvars \
dir --input examples --out conf
```
### Missing Key Behaviour
* --missingkey invalid (default)
```bash
echo "Hey! {{ .USER }} your home folder is {{ .MYHOME }}" | subvars
```
Output:
```
Hey! Jon your home folder is <no value>
```

* --missingkey zero
```bash
echo "Hey! {{ .USER }} your home folder is {{ .MYHOME }}" | subvars --missingkey zero
```
Output:
```
Hey! Jon your home folder is
```

* --missingkey error
```bash
echo "Hey! {{ .USER }} your home folder is {{ .MYHOME }}" | subvars --missingkey error
```
Output:
```
Hey! Jon your home folder is 2021/07/09 20:51:45 template: :1:40: 
executing "" at <.MYHOME>: map has no entry for key "MYHOME"
exit status 1
```

## Template Functions

In addition to the standard set of template actions and functions
that come with Go, `subvars` also incorporates sprig for additional, commonly used functions.

```bash
echo "Hey! {{ .USER | upper }} your home folder is {{ .HOME }}" | subvars
```
In the example, the username will be converted to upper case letters.
