# Charmil Config Package POC

## Steps to run:

1. Add your default configuration for the Host CLI in the [config.yaml](cmd/host/config.yaml) file
2. Run the Host CLI using `make run`

## End Result:

- #### [config.yaml](cmd/host/config.yaml) [Before running the Host CLI]:
  ```yaml
  key1: val1
  key2: val2
  key3: val3
  ```
- #### [config.yaml](cmd/host/config.yaml) [After running the Host CLI]:
  ```yaml
  key1: val1
  key2: val2
  key3: val3
  key4: val4
  plugins:
    pluginA:
      key5: val5
      key6: val6
      key7: val7
      key8: val8
  ```
