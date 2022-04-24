# embed

Small utility to create uint8_t array definition from a file.

## Example

If hello.txt file content is:

```txt
hello world
```

then running:

```shell
$ embed -in hello.txt -var file
```

prints to stdout the following:

```c++
const uint8_t file[] = {
  0x68,0x65,0x6c,0x6c,0x6f,0x20,0x77,0x6f,
  0x72,0x6c,0x64,
};

```