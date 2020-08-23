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
    breakout: true
    requirements <
      category: SERVER
      properties: "stationone"
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
      properties: "datastore"
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
    breakout: true
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
    sudo: true
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
nintents <
  job <
    name: "beertracker"
    go_path: "github.com/brotherlogic/beertracker"
  >  
  redundancy: GLOBAL
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
    name: "githubreceiver"
    go_path: "github.com/brotherlogic/githubreceiver"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "pullrequester"
    go_path: "github.com/brotherlogic/pullrequester"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "reminders"
    go_path: "github.com/brotherlogic/reminders"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "recordgetter"
    go_path: "github.com/brotherlogic/recordgetter"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "recordsorganiser"
    go_path: "github.com/brotherlogic/recordsorganiser"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "basicjob"
    go_path: "github.com/brotherlogic/basicjob"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "recordadder"
    go_path: "github.com/brotherlogic/recordadder"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "recordbudget"
    go_path: "github.com/brotherlogic/recordbudget"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "recordalerting"
    go_path: "github.com/brotherlogic/recordalerting"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "recordmatcher"
    go_path: "github.com/brotherlogic/recordmatcher"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "recordmover"
    go_path: "github.com/brotherlogic/recordmover"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "recordprinter"
    go_path: "github.com/brotherlogic/recordprinter"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "recordprocess"
    go_path: "github.com/brotherlogic/recordprocess"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "recordsales"
    go_path: "github.com/brotherlogic/recordsales"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "recordstats"
    go_path: "github.com/brotherlogic/recordstats"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "cdprocessor"
    go_path: "github.com/brotherlogic/cdprocessor"
    requirements <
      category: DISK
      properties: "raid"
    >
  >
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "recordwants"
    go_path: "github.com/brotherlogic/recordwants"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "etckiller"
    go_path: "github.com/brotherlogic/etckiller"
  >  
  redundancy: GLOBAL
>
nintents <
  job <
    name: "githubtasks"
    go_path: "github.com/brotherlogic/githubtasks"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "digitalwantlits"
    go_path: "github.com/brotherlogic/digitalwantlist"
  >  
  redundancy: REDUNDANT
>
