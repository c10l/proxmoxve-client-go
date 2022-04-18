# proxmoxve-client-go

WIP Go client for ProxMox Virtual Environment's JSON API.

This will be used in a Terraform provider so from this point on I will prioritise those APIs that I need. Contributions are welcome.

Instructions for running tests will be added in the future.

Current progress:

    ( ) access
    ( ) cluster
    ( ) nodes
        [ ] get
        ( ) {node}
            [ ] get
            ( ) apt
                [ ] get
                ( ) changelog
                    [ ] get
                ( ) repositories
                    [ ] get
                    [ ] post
                    [ ] put
                ( ) update
                    [ ] get
                    [ ] post
                ( ) versions
                    [ ] get
                ( ) capabilities
                    [ ] get
                    ( ) qemu
                        [ ] get
                        ( ) cpu
                            [ ] get
                        ( ) machines
                            [ ] get
                ( ) ceph
                ( ) certificates
                ( ) disks
                ( ) firewall
                ( ) hardware
                ( ) lxc
                ( ) network
                ( ) qemu
                ( ) replication
                ( ) scan
                ( ) sdn
                ( ) services
                ( ) storage
                ( ) tasks
                ( ) vzdump
                ( ) appinfo
                ( ) config
                ( ) dns
                ( ) execute
                ( ) hosts
                ( ) journal
                ( ) migrateall
                ( ) netstat
                ( ) query-url-metadata
                ( ) report
                ( ) rrd
                ( ) rrddata
                ( ) spiceshell
                ( ) startall
                ( ) status
                ( ) stopall
                ( ) subscription
                ( ) syslog
                ( ) termproxy
                ( ) time
                ( ) version
                ( ) vncshell
                ( ) vncwebsocket
                ( ) wakeonlan
    (o) pools
        [/] get
        [/] post
        (o) {poolid}
            [/] get
            [/] put
            [/] delete
    (o) storage
        [/] get
        [/] post
        (o) {storage}
            [/] get
            [/] put
            [/] delete
    (o) version
        [/] get
