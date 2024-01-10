# [Canonical Multipass](https://multipass.run/) provider for [DevPod](https://github.com/loft-sh/devpod)

[![Open in DevPod!](https://devpod.sh/assets/open-in-devpod.svg)](https://devpod.sh/open#https://github.com/minhio/devpod-provider-multipass)

## Prerequisites

- Install [Multipass](https://multipass.run/install)
- Install [DevPod](https://github.com/loft-sh/devpod)

## Getting started

The provider is available for auto-installation using DevPod CLI

```sh
devpod provider add minhio/devpod-provider-multipass
devpod provider use minhio/devpod-provider-multipass
```

Or the desktop app

![desktop-app-add-provider](.github/assets/desktop-app-add-provider.gif)

## Customize the Multipass Instance

This provides has the seguent options

| NAME                | REQUIRED | DESCRIPTION                                                     | DEFAULT   |
|---------------------|----------|-----------------------------------------------------------------|-----------|
| MULTIPASS_PATH      | true     | Path to multipass binary.                                       | multipass |
| MULTIPASS_IMAGE     | true     | Image to launch.                                                | lts       |
| MULTIPASS_CPUS      | true     | Number of CPUs to allocate.                                     | 2         |
| MULTIPASS_DISK_SIZE | true     | Disk space to allocate.                                         | 40G       |
| MULTIPASS_MEMORY    | true     | Amount of memory to allocate.                                   | 2G        |
| MULTIPASS_MOUNTS    | false    | Comma separated list of mounts in the format of "source:target" |           |
