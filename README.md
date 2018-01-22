# exegraph

Executable step by step graph processor

## Example

```go
  // Load document by HTTP
  loader, _ := s.graph.SetItem("loader", loader{})
  
  // Get links from document
  links, _ := s.graph.SetItem("links", exegraph.ExecuterFnk(s.processLinks), loader)
  
  // Process link and grab information
	s.graph.SetItem("process", exegraph.ExecuterFnk(s.processDocument), loader)
```
