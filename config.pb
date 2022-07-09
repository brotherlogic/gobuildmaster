nintents <
  job <
    name: "buildserver"
    go_path: "github.com/brotherlogic/buildserver"
    partial_bootstrap: true
    breakout: true
    requirements <
      category: DISK
      properties: "scratch"
    >
    requirements <
      category: HOST_TYPE
      properties: "Pi 4"
  >
  redundancy: REDUNDANT
  redundancy64: REDUNDANT
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
      properties: "clust2"
    >
  >
  redundancy: GLOBAL
>
nintents <
  job <
    name: "floppy"
    go_path: "github.com/brotherlogic/floppy"
    breakout: true
    requirements <
      category: SERVER
      properties: "floppy"
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
    name: "dstore"
    go_path: "github.com/brotherlogic/dstore"
    partial_bootstrap: true
    breakout: true
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
    name: "healthchecker"
    go_path: "github.com/brotherlogic/healthchecker"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "adventserver"
    go_path: "github.com/brotherlogic/adventserver"
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
    breakout: true
    requirements <
      category: DISK
      properties: "scratch"
    >
  >
  redundancy: GLOBAL
  not_on_server: "rdisplay"	
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
    partial_bootstrap: true
    sudo: true
  >  
  redundancy: GLOBAL
>
nintents <
  job <
    name: "printer"
    go_path: "github.com/brotherlogic/printer"
    breakout: true
    requirements <
      category: RECEIPT_PRINTER
    >
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "display"
    go_path: "github.com/brotherlogic/display"
    breakout: true
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "beertracker"
    go_path: "github.com/brotherlogic/beertracker"
    breakout: true
    requirements <	
      category: SERVER
      properties: "zero2"
    >
  >  
  count : 1
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
    name: "githubtasks"
    go_path: "github.com/brotherlogic/githubtasks"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "digitalwantlist"
    go_path: "github.com/brotherlogic/digitalwantlist"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "recordscores"
    go_path: "github.com/brotherlogic/recordscores"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "orgprinter"
    go_path: "github.com/brotherlogic/orgprinter"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "login"
    go_path: "github.com/brotherlogic/login"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "temp"
    go_path: "github.com/brotherlogic/temp"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "keymapper"
    go_path: "github.com/brotherlogic/keymapper"
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
    name: "backup"
    go_path: "github.com/brotherlogic/backup"
    requirements <
      category: DISK
      properties: "raid"
    >
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "recordvalidator"
    go_path: "github.com/brotherlogic/recordvalidator"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "printbridge"
    go_path: "github.com/brotherlogic/printbridge"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "stobridge"
    go_path: "github.com/brotherlogic/stobridge"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "lock"
    go_path: "github.com/brotherlogic/lock"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "queue"
    go_path: "github.com/brotherlogic/queue"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "recordfanout"
    go_path: "github.com/brotherlogic/recordfanout"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "recordcleaner"
    go_path: "github.com/brotherlogic/recordcleaner"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "euler"
    go_path: "github.com/brotherlogic/euler"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "recordsorg"
    go_path: "github.com/brotherlogic/recordsorg"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "dashboard"
    go_path: "github.com/brotherlogic/dashboard"
    breakout: true
    requirements <
      category: SERVER
      properties: "personal"
    >
  >  
  redundancy: GLOBAL
>
nintents <
  job <
    name: "secureproxy"
    go_path: "github.com/brotherlogic/secureproxy"
    breakout: true
    requirements <
      category: EXTERNAL
      properties: "external_ready"
    >
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "bandcampserver"
    go_path: "github.com/brotherlogic/bandcampserver"
  >  
  redundancy: REDUNDANT
>
nintents <
  job <
    name: "builder"
    go_path: "github.com/brotherlogic/builder"
  >  
  redundancy: REDUNDANT
>