# decode-uri

A CLI URI decoder for ENV variables.
Designed to split out `DATABASE_URL` style env variables into discrete parts.

## Example

```
$ export DATABASE_URL=postgres://user:pass@host.com:5432/path?k=v#f
$ decode-uri DATABASE_URL
export DATABASE_SCHEME=postgres
export DATABASE_USERNAME=user
export DATABASE_PASSWORD=pass
export DATABASE_HOST=host.com
export DATABASE_PORT=5432
export DATABASE_PATH=/path
export DATABASE_FRAGMENT=f
export DATABASE_QUERY_STRING=k=v
```

Parsed URI elements are returned formatted as export statements, if they exist,  prefixed with the value of the input env string before the first underscore.

### Usage

To get the parsed URI into your ENV you can do something like this:

```
$ eval "$(decode-uri DATABASE_URL)"
```

This command with exit with a non-zero exit status if a URI can't be parsed.
