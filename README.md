# kubectl-viz
A kubectl plugin that provides a visual output for `get` commands

[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

## Motivation

The motivation for this plugin is to provide a *simple* visual output for the `get` commands of `kubectl`, without creating significant overhead or anything that "gets in the way". The use of this plugin would ideally be as approachable as the use of the bare CLI tool. Specifically, the aim here is to avoid rendering a separate image file, require a web browser for viewing outputs, etc. 

While various use cases will be explored during development, the most relatable use case driving this plugin is to provide a visual output for a call such as `kubectl get pod --watch` so that the operator can watch the status of the cluster with a visual output.

## License

This work is licensed under the [MIT](https://opensource.org/licenses/MIT) open source license.

