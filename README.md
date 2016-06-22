## MemTool

Easier way to run command with memcache to replace telnet.

### Support:

1. get
2. set
3. del
4. flush

### Usage:

    NAME:
       MemcacheTool - Easier way to communicate with memcache
    
    USAGE:
       memtool [global options] command [command options] [arguments...]
       
    VERSION:
       20151010
       
    AUTHOR(S):
       Chen.Zhidong 
       
    COMMANDS:
         get        get value of key
         set        set value for key
         del        delete key
         flush      flush all keys
    
    GLOBAL OPTIONS:
       --host value         Host to connect (default: "127.0.0.1")
       --port value         Port to connect (default: "11211")
       --help, -h           show help
       --version, -v        print the version
       
    COPYRIGHT:
       http://sillydong.com

### Example:

1. get
    
    ./memtool get a
    ./memtool --host=127.0.0.1 --port=11211 get a
    
2. set

    ./memtool set a b
    ./memtool --host=127.0.0.1 --port=11211 set a b
    
3. del

    ./memtool del a
    ./memtool --host=127.0.0.1 --port=11211 del a
    
4. flush

    ./memtool flush
    ./memtool --host=127.0.0.1 --port=11211 flush
    
