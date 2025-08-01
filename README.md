[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/chx/ddev-kafka/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/chx/ddev-kafka/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/chx/ddev-kafka)](https://github.com/chx/ddev-kafka/commits)
[![release](https://img.shields.io/github/v/release/chx/ddev-kafka)](https://github.com/chx/ddev-kafka/releases/latest)

# DDEV Kafka

## Overview

This add-on integrates Kafka into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get chx/ddev-kafka
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev describe` | View service status and used ports for Kafka |
| `ddev logs -s kafka` | Check Kafka logs |

## Advanced Customization

To change the Docker image:

```bash
ddev dotenv set .ddev/.env.kafka --kafka-docker-image="busybox:stable"
ddev add-on get chx/ddev-kafka
ddev restart
```

Make sure to commit the `.ddev/.env.kafka` file to version control.

All customization options (use with caution):

| Variable | Flag | Default |
| -------- | ---- | ------- |
| `KAFKA_DOCKER_IMAGE` | `--kafka-docker-image` | `busybox:stable` |

## Credits

**Contributed and maintained by [@chx](https://github.com/chx)**
