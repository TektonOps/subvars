# Usage Examples


Several files are available in `examples` folder you can cd into the folder and test the following cases.


```bash
CLASS_NAME=nginx subvars < ingress-class.yaml
```

```bash
ALERT_THRESHOLD=2 subvars < prometheus_alert.conf
```

```bash
subvars < file.txt > file2.txt
```

```bash
subvars dir --input examples --out newdir
```

## Template Functions

In addition to the standard set of template actions and functions
that come with Go, `subvars` also incorporates sprig for additional, commonly used functions.

```bash
echo "Hey! {{ .USER | upper }} your home folder is {{ .HOME }}" | subvars
```
In the example, the username will be converted to upper case letters.
