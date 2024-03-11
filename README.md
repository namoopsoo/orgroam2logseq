## Migrate from orgroam to logseq markdown style
I have loved and used https://www.orgroam.com/ for 2 years but I wanted to take advantage of some of the many additional features of https://logseq.com/ and so this mini go utility is the result. 

This is written in Go just because I've been learning Go recently, but also Go is pretty fast.

## Example run 
The below example will take the example `example_org_roam` dir, and create `target_logseq_dir`, `target_logseq_dir/journals`, `target_logseq_dir/pages` and migrate journals and pages. 

Currently `example_org_roam/assets` are not moved/copied . But that's just a simple copy command away. 

```sh
go run example_org_roam target_logseq_dir
```
