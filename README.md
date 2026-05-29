# flint

Field linter and lookup tool.

## Installation

```shell
go install .
```

## Building

```shell
go build .
```

## Usage

```shell
flint resolve \
  -p beats:/path/to/beats \
  -p integrations:/path/to/integrations \
  -t detectionrules
  /path/to/detection-rules/rules/some_rule.toml
```
