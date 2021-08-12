# LEB128

Little Endian Base 128 | Variable-Length Code Compression

## Usage

```
bs, _ := leb128.EncodeUnsigned(big.NewInt(200000))
// [192 154 12]
bi, _ := leb128.DecodeUnsigned(bytes.NewReader(bs))
// 200000
```

## Uses

- LEB128 is used in the [WebAssembly](https://webassembly.github.io/spec/core/binary/values.html#integers) binary
  encoding for all integer literals.
- LEB128 is used by [Candid](https://github.com/dfinity/candid/blob/master/spec/Candid.md#serialisation) to serialise
  types into a binary representation for transfer between services.
