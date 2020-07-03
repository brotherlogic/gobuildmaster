nintents <
  job <
    name: "buildserver"
    go_path: "github.com/brotherlogic/buildserver"
    partial_bootstrap: true
    requirements <
      category: DISK
      properties: "scratch"
    >
  >
  redundancy: REDUNDANT
>
nintents <
  job <
    go_path: "github.com/brotherlogic/filecopier"
    name: "filecopier"
    partial_bootstrap: true
    breakout: true
  >
  redundancy: GLOBAL
  no_master: true
>
nintents <
  job <
    name: "versiontracker"
    go_path: "github.com/brotherlogic/versiontracker"
    partial_bootstrap: true
    breakout: true
  >
  redundancy: GLOBAL
  no_master: true
>
nintents <
  job <
    name: "keystore"
    go_path: "github.com/brotherlogic/keystore"
    requirements <
      category: SERVER
      properties: "runner"
    >
  >
  redundancy: GLOBAL
>
nintents <
  job <
    name: "datastore"
    go_path: "github.com/brotherlogic/datastore"
    requirements <
      category: DISK
      properties: "keystore"
    >
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
    name: "dropboxsync"
    go_path: "github.com/brotherlogic/dropboxsync"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "frametracker"
    go_path: "github.com/brotherlogic/frametracker"
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
    name: "logging"
    go_path: "github.com/brotherlogic/logging"
    requirements <
      category: DISK
      properties: "scratch"
    >
  >
  redundancy: REDUNDANT	
>
nintents <
  job <
    name: "beerserver"
    go_path: "github.com/brotherlogic/beerserver"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "executor"
    go_path: "github.com/brotherlogic/executor"
    breakout: true
  >  
  redundancy: GLOBAL
>
nintents <
  job <
    name: "provisioner"
    go_path: "github.com/brotherlogic/provisioner"
    breakout: true
  >  
  redundancy: GLOBAL
>
nintents <
  job <
    name: "printer"
    go_path: "github.com/brotherlogic/printer"
    partial_bootstrap: true
    requirements <
      category: SERVER
      properties: "printer"
    >
  >  
  count: 1
>