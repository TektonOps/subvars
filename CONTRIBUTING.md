# Contributing to SubVars

We want to make contributing to this project as easy and transparent as
possible.

## Project structure
- `cmd` - Contains the file and helpers package.
  - `dir` Content related to folder parsing.
  - `helpers` Helper functions.
  
### Documentation

- `docs` - Contains the documentation in Markdown format.
  - `home.md` Is the main page rendered when docs website is loaded.
  - `install.md` Contains the installation instructions for different packages.

### Checking Locally

Mkdocs is used for documentation rendering from markdown, you can download
the cli and test locally before opening a pull request.

### Install

```bash
pip install mkdocs-material
pip install mdx_include
pip install mkdocs-minify-plugin
```

### Serve locally

```bash
mkdocs serve
```

## Commits

Commit messages should be well formatted, and to make that "standardized", we
are using Conventional Commits.

```shell

  <type>[<scope>]: <short summary>
     │      │             │
     │      │             └─> Summary in present tense. Not capitalized. No
     |      |                 period at the end. 
     │      │
     │      └─> Scope (optional): eg. common, compiler, authentication, core
     │
     └─> Type: chore, docs, feat, fix, refactor, style, or test.
     
```

You can follow the documentation on
[their website](https://www.conventionalcommits.org).

## Pull Requests

We actively welcome your pull requests.

1. Fork the repo and create your branch from `master`.
2. If you've added code that should be tested, add tests.
3. If you've changed APIs, update the documentation.
4. Make sure your code lints (`make lint`).
5. Make sure your code is well formatted (`make fmt`).

## Issues

We use GitHub issues to track public bugs. Please ensure your description is
clear and has sufficient instructions to be able to reproduce the issue.

## License

By contributing to PingMe, you agree that your contributions will be licensed
under the LICENSE file in the root directory of this source tree.
