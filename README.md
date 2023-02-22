# Go example for [Dagger](https://dagger.io/)

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/sagikazarmark/dagger-go-example/ci.yaml?style=flat-square)](https://github.com/sagikazarmark/dagger-go-example/actions?query=workflow%3ACI)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/sagikazarmark/dagger-go-example/dagger.yaml?style=flat-square)](https://github.com/sagikazarmark/dagger-go-example/actions?query=workflow%3ADagger)
[![Codecov](https://img.shields.io/codecov/c/github/sagikazarmark/dagger-go-example?style=flat-square)](https://codecov.io/gh/sagikazarmark/dagger-go-example)

This repository serves as an example for using [Dagger](https://dagger.io/) Go SDK.

## Usage

Run tests:

```shell
mage -d ci -w . test
```

Run linter:

```shell
mage -d ci -w . lint
```

## License

The project is licensed under the [MIT License](LICENSE).
