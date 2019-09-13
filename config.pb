nintents <
  job <
    name: "buildserver"
    go_path: "github.com/brotherlogic/buildserver"
    bootstrap: true
    requirements <
      category: DISK
      properties: "scratch"
    >
    requirements <
      category: SERVER
      properties: "framethree"
    >
  >
  redundancy: REDUNDANT
>
nintents <
  job <
    go_path: "github.com/brotherlogic/filecopier"
    name: "filecopier"
    bootstrap: true
    breakout: true
  >
  redundancy: GLOBAL
>
nintents <
  job <
    go_path: "github.com/brotherlogic/githubcard"
    name: "githubcard"
   >
   redundancy: REDUNDANT
>
nintents <
  job <
    name: "keystore"
    go_path: "github.com/brotherlogic/keystore"
    requirements <
      category: DISK
      properties: "keystore"
    >
  >
  redundancy: GLOBAL
>
nintents <
  job <
    name: "monitor"
    go_path: "github.com/brotherlogic/monitor"
    requirements <
      category: DISK
      properties: "scratch"
    >
  >
  redundancy: REDUNDANT	
>
nintents <
  job <
    name: "versionserver"
    go_path: "github.com/brotherlogic/versionserver"
    requirements <
      category: SERVER
      properties: "runner"
    >
  >
  count: 1
>
nintents <
  job <
    name: "recordcollection"
    go_path: "github.com/brotherlogic/recordcollection"
  >
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "proxy"
    go_path: "github.com/brotherlogic/proxy"
    requirements <
      category: EXTERNAL
      properties: "external_ready"
    >
  >
  count: 1
>
nintents <
  job <
    name: "githubreceiver"
    go_path: "github.com/brotherlogic/githubreceiver"
  >  
  count: REDUNDANT
>
nintents <
  job <
    name: "recordgetter"
    go_path: "github.com/brotherlogic/recordgetter"
  >
  redundancy: REDUNDANT
>
