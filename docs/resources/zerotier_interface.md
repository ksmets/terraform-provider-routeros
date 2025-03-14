# routeros_zerotier_interface (Resource)


## Example Usage
```terraform
resource "zerotier_identity" "identity" {}

resource "zerotier_network" "network" {
  name = "test"
}

resource "zerotier_member" "member" {
  authorized              = true
  member_id               = zerotier_identity.identity.id
  name                    = "test"
  network_id              = zerotier_network.network.id
  hidden                  = false
  allow_ethernet_bridging = true
  no_auto_assign_ips      = true
}

resource "routeros_zerotier" "zt1" {
  comment    = "ZeroTier Central"
  identity   = zerotier_identity.identity.private_key
  interfaces = ["all"]
  name       = "zt1"
}

resource "routeros_zerotier_interface" "zerotier1" {
  allow_default = false
  allow_global  = false
  allow_managed = false
  instance      = routeros_zerotier.zt1.name
  name          = "zerotier1"
  network       = zerotier_network.network.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `instance` (String) The ZeroTier instance name.
- `name` (String) Name of the ZeroTier interface.
- `network` (String) The ZeroTier network identifier.

### Optional

- `allow_default` (Boolean) An option to override the default route.
- `allow_global` (Boolean) An option to allow overlapping public IP space by the ZeroTier routes. .
- `allow_managed` (Boolean) An option to allow assignment of managed IPs.
- `arp_timeout` (String) ARP timeout is time how long ARP record is kept in ARP table after no packets are received from IP. Value auto equals to the value of arp-timeout in IP/Settings, default is 30s. Can use postfix `ms`, `s`, `m`, `h`, `d` for milliseconds, seconds, minutes, hours or days. If no postfix is set then seconds (s) is used.
- `comment` (String)
- `disable_running_check` (Boolean) An option to force the `running` property to true.
- `disabled` (Boolean)

### Read-Only

- `bridge` (Boolean) A flag whether the ZeroTier interface is bridged.
- `dhcp` (Boolean) A flag whether the ZeroTier interface obtained an IP address.
- `id` (String) The ID of this resource.
- `mac_address` (String) Current mac address.
- `mtu` (Number) Layer2 Maximum transmission unit. [See](https://wiki.mikrotik.com/wiki/Maximum_Transmission_Unit_on_RouterBoards).
- `network_name` (String) The ZeroTier network name.
- `running` (Boolean)
- `status` (String) The status of the ZeroTier connection.
- `type` (String) The ZeroTier network type.

## Import
Import is supported using the following syntax:
```shell
#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/zerotier/interface get [print show-ids]]
terraform import routeros_zerotier_interface.zerotier1 '*1'
```
