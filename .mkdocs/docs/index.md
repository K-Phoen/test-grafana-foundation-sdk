---
title: Home
hide:
  - navigation
---

# Grafana Foundation SDK

A set of libraries for manipulating and generating Grafana resources
– dashboards, alerts, … – as-code, in various languages.

**_Types_**, **_builder libraries_** and **JSON to code converters** are
provided for a range of versions of Grafana.

!!! note

    The various SDKs are generated by [`cog`][cog] from
    [schemas exposed by Grafana][kind-registry].

## Maturity

The Grafana Foundation SDK should be considered as "public preview". While it is used by Grafana Labs in production, it still is under active development.

Additional information can be found in [Release life cycle for Grafana Labs](https://grafana.com/docs/release-life-cycle/).

!!! note

    Bugs and issues are handled solely by Engineering teams. On-call support or SLAs are not available.

## License

[Apache 2.0 License](https://www.apache.org/licenses/LICENSE-2.0)

[cog]: <https://github.com/grafana/cog>
[kind-registry]: <https://github.com/grafana/kind-registry>
