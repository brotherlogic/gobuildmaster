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
    name: "stest"
    go_path: "github.com/brotherlogic/stest"
  >  
  redundancy: REDUNDANT
>
