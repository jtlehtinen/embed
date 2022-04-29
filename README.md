# embed

Small utility to embed a file contents into a  uint8_t array.

## Example

If hello.txt file content is:

```txt
hello world
```

then running:

```shell
$ embed -var bytes hello.txt
```

prints to stdout the following:

```c++
const uint8_t bytes[] = {
  0x68,0x65,0x6c,0x6c,0x6f,0x20,0x77,0x6f,
  0x72,0x6c,0x64,
};
```