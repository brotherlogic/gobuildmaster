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
  count: 1
>
nintents <
  job <
    go_path: "github.com/brotherlogic/beerserver"
    name: "beerserver"
  >
 count: 1
>
nintents <
  job <
    go_path: "github.com/brotherlogic/cardserver"
    name: "cardserver"
  >
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "recordgetter"
    go_path: "github.com/brotherlogic/recordgetter"
    non_bootstrap: true
  >
  count: 1
>
intents <
  spec <
    name: "github.com/brotherlogic/monitor"
  >
>
nintents <
  job <
    go_path: "github.com/brotherlogic/recordsorganiser"
    name: "recordsorganiser"
    non_bootstramp: true
   >
  count: 1
>
nintents <
  job <
    go_path: "github.com/brotherlogic/githubcard"
    name: "githubcard"
    non_bootstrap: true
   >
  count: 1
>
nintents <
  job <
    go_path: "github.com/brotherlogic/reminders"
    name: "reminders"
    non_bootstrap: true
   >
  count: 1
>
nintents <
  job <
    go_path: "github.com/brotherlogic/recordprocess"
    name: "recordprocess"
   >
  count: 1
>
nintents <
  job <
    go_path: "github.com/brotherlogic/recordmover"
    name: "recordmover"
    non_bootstrap: true
   >
  count: 1
>
intents <
  spec <
    name: "github.com/brotherlogic/cdprocessor"
    cds: true
  >
  count: 3
>
nintents <
  job <
    go_path: "github.com/brotherlogic/recordwants"
    name: "recordwants"
  >
  count: 1
>
nintents <
  job <
    go_path: "github.com/brotherlogic/recordalerting"
    name: "recordalerting"
    non_bootstrap: true
  >
  count: 1
>
nintents <
  job <
    go_path: "github.com/brotherlogic/filecopier"
    name: "filecopier"
  >
  redundancy: GLOBAL
>
nintents <
  job <
    name: "tracer"
    go_path: "github.com/brotherlogic/tracer"
  >
  count: 1
>
nintents <
  job <
    name: "monitor"
    go_path: "github.com/brotherlogic/monitor"
  >
  redundancy: REDUNDANT	
>
nintents <
  job <
    name: "buildserver"
    go_path: "github.com/brotherlogic/buildserver"
    requirements <
      category: DISK
      properties: "scratch"
    >
  >
  count : 1
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
    name: "location"
    go_path: "github.com/brotherlogic/location"
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
    name: "led"
    go_path: "github.com/brotherlogic/led"
    sudo: true
    requirements <
      category: SERVER
      properties: "leddisplay"
    >
  >  
  count: 1
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
    name: "recordprinter"
    go_path: "github.com/brotherlogic/recordprinter"
    non_bootstrap: true
  >  
  count: 1
>
nintents <
  job <
    name: "alerter"
    go_path: "github.com/brotherlogic/alerter"
    non_bootstrap: true
  >  
  count: 1
>
nintents <
  job <
    name: "recordsales"
    go_path: "github.com/brotherlogic/recordsales"
    non_bootstrap: true
  >  
  count: 1
>
nintents <
  job <
    name: "datacollector"
    go_path: "github.com/brotherlogic/datacollector"
    non_bootstrap: true
  >  
  count: 1
>
nintents <
  job <
    name: "dataviewer"
    go_path: "github.com/brotherlogic/dataviewer"
    non_bootstrap: true
  >  
  count: 1
>