# Identicon

A port of GitHub's [identicon](https://github.blog/2013-08-14-identicons/)
algorithm to Go, based on
[the Rust implementation](https://github.com/dgraham/identicon) (licensed
under the
[MIT License](https://github.com/dgraham/identicon/blob/master/LICENSE)).

![Identicons](https://github.blog/wp-content/uploads/2013/08/a3c4e2a0-04df-11e3-824c-7378e6550707.png)

## To Build

To build the binary, run:

```bash
go build ./cmd/identicon
```

## To Use

To generate the identicon for your GitHub account, run:

```bash
echo -n <your-user-id> | /path/to/identicon > /path/to/output.png
```

You may check out your GitHub user ID at
`https://api.github.com/users/<your-username>`. Verify that you get the same
result as from `https://github.com/identicons/<your-username>.png`.

To generate a random identicon, simply pass in any random input to the
program.

## License

Copyright (c) 2022 Zhong Ruoyu. Licensed under the [MIT License](LICENSE).
