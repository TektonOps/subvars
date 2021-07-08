
<h2 align="center">
  <p align="center"><img width=50% src="assets/index_img.png"></p>
    <br>
</h2>

# Welcome!

Substitute Variables (subvars) is a small utility which provides a way to render any [Go templates](https://golang.org/pkg/text/template/) from command line recognizing the object being passed in and drawing attributes from the object to create wanted text. It is very useful for template driven configuration files.

It uses [sprig v3](https://github.com/Masterminds/sprig) for [template functions](https://masterminds.github.io/sprig) which provides additional functions apart from standard library.

## Reading and Rednering

```bash
subvars [global options] command [command options] [arguments...]
```

* Directory `dir` subcommand lets you render all files in a folder & subfolder, and writes the output to `stdout`.
  you can also set an output folder where rendered files will be saved insted of stdout with `--outdir` flag. The filename will be same. If the folder does not exists it will be created automatically.
* `subvars` reads the template directly from `stdin`
* Renderd output will be written to `stdout`

See the Usage example for details.