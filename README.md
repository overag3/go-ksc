# go-ksc #
go-ksc is a Go client library for accessing the KSC (Kaspersky) Open API.

## Usage ##

```go
import (
    "github.com/pixfid/go-ksc/kaspersky"
)

```

Construct a new KSC client, then use the various services on the client to
access different parts of the KSC (Kaspersky) Open API. For example:

```go
        ctx := context.Background()
    	cfg := kaspersky.Config {
    		Username: "login",
    		Password: "password",
    		Server: fmt.Sprintf(`https://%s:%s`, "ip", "port"),
    	}
    
    	client := kaspersky.New(cfg)
    	client.KSCAuth(ctx)

        //Get List of Windows domain in the network.
        groups_domains, err := client.HostGroup.GetDomains(context.Background())
```

As example find online hosts:
```go
func Online(ctx context.Context, client *kaspersky.Client) *FullHostsInfo {
	hField := config.Config.HParams
	chunks := &FullHostsInfo{}
	hostsParam := kaspersky.HGParams{
		WstrFilter: `
		(&
			(KLHST_WKS_GROUPID_GP <> 4)
			(KLHST_WKS_STATUS&1<>0)
		)`,
		VecFieldsToReturn: hField,
		PParams: kaspersky.PParams{
			KlsrvhSlaveRecDepth:    0,
			KlgrpFindFromCurVsOnly: true,
		},
		LMaxLifeTime: 100,
	}

	accessor, _, _ := client.HostGroup.FindHosts(ctx, hostsParam)
	count, _, _ := client.ChunkAccessor.GetItemsCount(ctx, accessor.StrAccessor)
	_, _ = client.ChunkAccessor.GetItemsChunk(ctx, kaspersky.ItemsChunkParams{
		StrAccessor: accessor.StrAccessor,
		NStart:      0,
		NCount:      count.Int,
	}, chunks)

	client.ChunkAccessor.Release(ctx, accessor.StrAccessor)
	return chunks
}
```


NOTE: Using the [context](https://godoc.org/context) package, one can easily
pass cancelation signals and deadlines to various services of the client for
handling a request. In case there is no context available, then `context.Background()`
can be used as a starting point.


## License ##

This library is distributed under the  GNU GENERAL PUBLIC LICENSE Version 3 found in the [LICENSE](./LICENSE)
file.