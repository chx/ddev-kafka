[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/chx/ddev-apache-kafka/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/chx/ddev-apache-kafka/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/chx/ddev-apache-kafka)](https://github.com/chx/ddev-apache-kafka/commits)
[![release](https://img.shields.io/github/v/release/chx/ddev-apache-kafka)](https://github.com/chx/ddev-apache-kafka/releases/latest)

# DDEV Apache Kafka

## Overview

This add-on integrates Apache Kafka into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get chx/ddev-apache-kafka
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev describe` | View service status and used ports for Apache Kafka |
| `ddev logs -s apache-kafka` | Check Apache Kafka logs |

## Credits

**Contributed and maintained by [@chx](https://github.com/chx) with sponsorship from [Tag1 Consulting](https://tag1consulting.com/)**
