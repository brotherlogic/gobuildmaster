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
  count: 1
>
nintents <
  job <
    name: "recordgetter"
    go_path: "github.com/brotherlogic/recordgetter"
  >
  count: 1
>
intents <
  spec <
    name: "github.com/brotherlogic/monitor"
    external: true
  >
  count: 1
>
nintents <
  job <
    go_path: "github.com/brotherlogic/recordsorganiser"
    name: "recordsorganiser"
   >
  count: 1
>
nintents <
  job <
    go_path: "github.com/brotherlogic/githubcard"
    name: "githubcard"
   >
  count: 1
>
nintents <
  job <
    go_path: "github.com/brotherlogic/reminders"
    name: "reminders"
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
