nintents <
  job <
    name: "buildserver"
    go_path: "github.com/brotherlogic/buildserver"
    bootstrap: true
  >
  redundancy: REDUNDANT
>
nintents <
  job <
    go_path: "github.com/brotherlogic/filecopier"
    name: "filecopier"
    bootstrap: true
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
