# Configuration

```bash
subvars [global options] command [command options] [arguments...]
```

```bash
subvars dir [command options] [arguments...]
```

* Directory `dir` subcommand lets you render all files in a folder & subfolder, and writes the output to `stdout`.
  you can also set an output folder where rendered files will be saved insted of stdout with `--outdir` flag. The filename will be same. If the folder does not exists it will be created automatically.
* `subvars` reads the template directly from `stdin`
* Renderd output will be written to `stdout`

## Flags

### Missing Key

* If `missingkey` is unset the  default is value is set to `invalid`,
  and it follows the default behaviour of [the golang template library](https://golang.org/pkg/text/template/#Template.Option).

* Missing keys in the template will be substituted with the string `<no value>`.

* If `missingkey` is set to `zero`, missing keys will be substituted with zero value of data type (ie: an empty
  string).

* If `missingkey` is set to `error`, `subvars` will fail and
  returns an error to the caller when missing any key.

* Settings can be changed with `--missingkey` prefix or by exporting environment variable `SUBVARS_MISSINGKEY`.

### Prefix
Prefix flag `--prefix` will match a given prefix pattern of all env variables and render only matching the prefixIt can also be configured by exporting environment variable `SUBVARS_PREFIX`.


### Input
Input flag `--input` is available for the subcommand `dir`, when using subcommand you can specify an input folder containing the template files.

It can also be configured by exporting environment variable `SUBVARS_INPUTDIR`.


### Out
Input flag `--out` is available for the subcommand `dir`, when using subcommand you can specify an output folder where rendered files will be saved.

If the folder does not exist it will be created automatically, the output filename will remain the same as input templates.

This setting can also be configured by exporting environment variable `SUBVARS_OUTDIR`.

For more details please check the [examples](03-usage-examples.md) page.
