nintents <
  job <
    name: "keystore"
    go_path: "github.com/brotherlogic/keystore"
    requirements <
      category: DISK
      properties: "disk1"
    >
  >
  count: 1
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
intents <
  spec <
    name: "github.com/brotherlogic/beerserver"
  >
 count: 1
>
intents <
  spec <
    name: "github.com/brotherlogic/cardserver"
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
intents <
  spec <
    name: "github.com/brotherlogic/recordsorganiser"
   >
  count: 1
>
intents <
  spec <
    name: "github.com/brotherlogic/githubcard"
   >
  count: 1
>
intents <
  spec <
    name: "github.com/brotherlogic/reminders"
   >
  count: 1
>
intents <
  spec <
    name: "github.com/brotherlogic/recordcollection"
   >
  count: 1
>
intents <
  spec <
    name: "github.com/brotherlogic/recordprocess"
   >
  count: 1
>
intents <
  spec <
    name: "github.com/brotherlogic/recordmover"
   >
  count: 3
>
intents <
  spec <
    name: "github.com/brotherlogic/movelist"
   >
  count: 3
>
intents <
  spec <
    name: "github.com/brotherlogic/cdprocessor"
    cds: true
  >
  count: 3
>
intents <
  spec <
    name: "github.com/brotherlogic/recordwants"
  >
  count: 3
>
intents <
  spec <
    name: "github.com/brotherlogic/recordalerting"
  >
  count: 3
>
intents <
  spec <
    name: "github.com/brotherlogic/filecopier"
  >
  count: 1
>
