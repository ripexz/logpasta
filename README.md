# logpasta

Simple CLI for [Logpasta](https://www.logpasta.com), just run `your_command | logpasta`.

### Installation

- Via Go: `go get github.com/ripexz/logpasta`
- [Download binary](https://github.com/ripexz/logpasta/releases) and put it in your `$PATH`

### Usage

- `any_command | logpasta [flags]`
- `logpasta [flags] any string at all`
- `logpasta [flags] < your_file`

To see the current CLI version, run `logpasta version`

### Config

You can override the default configuration by passing flags:
- `-u` - `string` - override default logpasta URL (default: `https://www.logpasta.com`)
- `-s` - `bool` - suppress output, unless the request fails (default: `true`)
- `-d` - `bool` - output debug information (default: `false`)
- `-t` - `int` - http client timeout in seconds (default: `5`)

You can also set up environmental variables to make the job easier, however flags will override these:
- `LOGPASTA_URL` - `string` - see `-u` flag
- `LOGPASTA_SILENT` - `bool` - see `-s` flag
- `LOGPASTA_DEBUG` - `bool` - see `-d` flag
- `LOGPASTA_TIMEOUT` - `int` - see `-t` flag
