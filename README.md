# Sparsify files with contiguous zero regions

There are situations when you have a file with a lot of zeros, and you want to compress it. This is a simple utility that will compress the file by removing the contiguous zero regions.

It is especially useful for core dumps, memory dumps, and other files that contain a lot of zeros.

## Usage

```bash

$ sparsify --block-size <block_size> --input <input_file> --output <output_file>

```

* `block_size` - the size of the blocks that will be checked for zeros and sparsified
* `input_file` - the file that will be sparsified; it can be - for stdin
* `output_file` - the file where the sparsified data will be written

## Usage for core dumps

In linux you can pipe the core dump through sparsify to avoid writing the zeros to disk by setting the following value in `/proc/sys/kernel/core_pattern`:

```bash

|/path/to/sparsify --block-size 4096 --input - --output /path/to/output_file

```

Example:

```bash

echo "|/usr/local/bin/sparsify --block-size 4096 --input - --output /var/cores/core.%p.%e.%t" > /proc/sys/kernel/core_pattern

```