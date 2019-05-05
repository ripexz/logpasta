# logpasta

Simple CLI for [Logpasta](https://www.logpasta.com), just run `your_command | logpasta`.

### Usage

- `any_command | logpasta [flags]`
- `logpasta [flags] any string at all`

### Config

You can override the default configuration by passing flags:
- `-u` - `string` - override default logpasta URL (default: `https://www.logpasta.com`)
- `-s` - `bool` - suppress output, unless the request fails (default: `true`)
- `-d` - `bool` - output debug information (default: `false`)

You can also set up environmental variables to make the job easier, however flags will override these:
- `LOGPASTA_URL` - `string` - see `-u` flag
- `LOGPASTA_SILENT` - `bool` - see `-s` flag
- `LOGPASTA_DEBUG` - `bool` - see `-d` flag