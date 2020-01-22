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
    name: "alerter"
    go_path: "github.com/brotherlogic/alerter"
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
    name: "dropboxsync"
    go_path: "github.com/brotherlogic/dropboxsync"
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
    go_path: "github.com/brotherlogic/githubcard"
    name: "githubcard"
   >
   redundancy: REDUNDANT
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
    name: "executor"
    go_path: "github.com/brotherlogic/executor"
    breakout: true
  >  
  redundancy: GLOBAL
>
nintents <
  job <
    name: "tracer"
    go_path: "github.com/brotherlogic/tracer"
  >
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "printer"
    go_path: "github.com/brotherlogic/printer"
    requirements <
      category: SERVER
      properties: "printer"
    >
  >  
  count: 1
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
    go_path: "github.com/brotherlogic/beerserver"
    name: "beerserver"
  >
 redundancy: REDUNDANT
>
nintents <
  job <
    name: "keystorebackup"
    go_path: "github.com/brotherlogic/keystorebackup"
    requirements <
      category: DISK
      properties: "raid1"
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
    name: "recordadder"
    go_path: "github.com/brotherlogic/recordadder"
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
    go_path: "github.com/brotherlogic/recordwants"
    name: "recordwants"
  >
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "wantslist"
    go_path: "github.com/brotherlogic/wantslist"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    go_path: "github.com/brotherlogic/reminders"
    name: "reminders"
   >
   redundancy: REDUNDANT
>
nintents <
  job <
    go_path: "github.com/brotherlogic/recordsorganiser"
    name: "recordsorganiser"
   >
   redundancy: REDUNDANT
>
nintents <
  job <
    go_path: "github.com/brotherlogic/recordprocess"
    name: "recordprocess"
   >
     redundancy: REDUNDANT
>
nintents <
  job <
    go_path: "github.com/brotherlogic/recordmover"
    name: "recordmover"
   >
   redundancy: REDUNDANT	
>
nintents <
  job <
    go_path: "github.com/brotherlogic/recordalerting"
    name: "recordalerting"
  >
  redundancy: REDUNDANT
>
nintents <
  job <
    go_path: "github.com/brotherlogic/recordprinter"
    name: "recordprinter"
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
    name: "cdprocessor"
    go_path: "github.com/brotherlogic/cdprocessor"
    requirements <
      category: DISK
      properties: "music"
    >
  >
  redundancy: REDUNDANT
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
    name: "recordbudget"
    go_path: "github.com/brotherlogic/recordbudget"
  >  
  redundancy: REDUNDANT
>