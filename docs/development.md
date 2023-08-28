## Development Setup
Note: steps 3 & 4 assume a developer is on some distribution of Debian/Ubuntu. If not on that distribution,
you can simply look at what is happening in `taskfile.yaml` and replicate that in your environment.

1. Install [golang](https://go.dev/doc/install)
2. Install [taskfile.dev](https://taskfile.dev/installation/) on your machine.

3. Run `task install_deps` in the root of this repo.

4. Run `task run` in the root of this repo.

If all passes without errors, you should be set to use this binary. Checkout the `taskfile.dev` for additional features
commands you can run.

**Note**: You can read through the [first_time_setup.md](https://github.com/jmillerv/go-dj/blob/main/docs/first_time_setup.md) document for troubleshooting help if the taskfile doesn't work. 