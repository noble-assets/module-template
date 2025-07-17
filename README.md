# Module Template

## Required modifications

### proto/

- `generate.sh`:

  - Update the package and folder names: `template.dev`. `template`

- `buf.gen.pulsar.yaml`:

  - Update the package name: `template.dev`

- Update the folder name inside `proto/noble` to match the package name.

Inside this new folder, modify all the occurrencies of the word `template` in docstring and packages names.
