# github.com/Infrawrench/infrawrench-go

Generated Go client for the Infrawrench API.

**API version `1.19.0`.** A Go module takes its version from a VCS
tag rather than from a manifest field, so the API version lives in this README
and in the `APIVersion` constant — check that constant, not the module tag,
when you need to know which API shape you have.

**Do not edit this module by hand** — it is regenerated from `openapi.json` and
is not checked into the repository. Run
`pnpm --filter @infrawrench/web generate:sdk` to rebuild it; the generator lives
in [`app/packages/web/scripts/sdk`](https://github.com/Infrawrench/Infrawrench/tree/main/app/packages/web/scripts/sdk).

## Install

```sh
go get github.com/Infrawrench/infrawrench-go
```

Go 1.24 or newer. No dependencies: the module requires nothing but the
standard library, and `go.mod` has no `require` block at all.

## Usage

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Infrawrench/infrawrench-go"
)

func main() {
	ctx := context.Background()
	client := infrawrench.NewAPIV1Client(
		infrawrench.WithAPIKey(os.Getenv("INFRAWRENCH_API_KEY")),
		infrawrench.WithOrgID(os.Getenv("INFRAWRENCH_ORG_ID")),
	)

	accounts, err := client.Accounts.List(ctx, nil)
	if err != nil {
		var apiErr *infrawrench.APIError
		if errors.As(err, &apiErr) {
			fmt.Println(apiErr.StatusCode, apiErr.Code, string(apiErr.Body))
		}
		return
	}
	fmt.Println(len(accounts))
}
```

## Conventions

- **Context first.** Every call takes a `context.Context` as its first
  argument and a variadic `...RequestOption` as its last. Timeouts and
  cancellation belong on the context, so there is no per-call timeout option.
- **Dotted namespaces.** Calls mirror the URL structure, so
  `POST /api/org/{orgId}/accounts/{id}/sync` is `client.Accounts.Sync(...)`.
- **Parameters.** A call with at least one mandatory parameter takes its params
  struct by value; a call where everything is optional takes a pointer, so
  `nil` means "no arguments". A call with no parameters at all takes neither.
- **Organization id.** Pass `WithOrgID` once when constructing the
  client and every scoped call can leave `OrgID` unset; set
  `OrgID` on an individual call to override it. With neither, the call
  returns an error wrapping `ErrMissingPathParam` rather than sending a
  malformed URL.
- **Optional fields are pointers.** Go cannot tell an omitted `false` from a
  deliberate one, so anything the wire may omit or null is a `*T`. Slices and
  maps keep their own nil.
- **Errors are returned, never panicked.** Any non-2xx response comes back as
  `*APIError`, carrying `StatusCode`, the raw `Body`, the decoded `Data`
  and the machine-readable `Code` when the API sends one. Use `errors.As`, or
  the `AsAPIError` shorthand.
- **Downloads stream.** An endpoint that returns a file returns an
  `io.ReadCloser`; close it.

## Scope

This module covers the published API surface only. Operations marked
`x-internal` in the spec — the admin surface, webhook receivers, desktop sync,
push registration, and the browser auth redirects — are not generated, so there
is no namespace for them to be called through.

Topics: infrawrench, sdk, api-client, openapi, infrastructure, cloud, devops.

## License

MIT — see [`LICENSE`](./LICENSE). Copyright (c) 2026 Infrawrench LLC.

Note that this client is more permissively licensed than the service it talks
to: the Infrawrench source is BUSL-1.1, but the generated clients are MIT
so you can link one into your own software without inheriting those terms.

Maintained by Infrawrench LLC <astrid@infrawrench.com>. Documentation: <https://infrawrench.com/docs/team-and-billing/client-sdks>.
Issues: <https://github.com/Infrawrench/Infrawrench/issues>
