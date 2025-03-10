# routeros_ipv6_dhcp_server_option (Resource)


## Example Usage
```terraform
resource "routeros_ipv6_dhcp_server_option" "test" {
  name  = "domain-search"
  code  = 24
  value = "0x07'example'0x05'local'0x00"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `code` (Number) Dhcp option [code](https://www.ipamworldwide.com/ipam/isc-dhcpv6-options.html).
- `name` (String) Descriptive name of the option.

### Optional

- `comment` (String)
- `value` (String) Parameter's value. Available data types for options are:
    - `'test'` -> ASCII to Hex 0x74657374
    - `'10.10.10.10'` -> Unicode IP to Hex 0x0a0a0a0a
    - `s'10.10.10.10'` -> ASCII to Hex 0x31302e31302e31302e3130
    - `s'160'` -> ASCII to Hex 0x313630
    - `'10'` -> Decimal to Hex 0x0a
    - `0x0a0a` -> No conversion
    - `$(VARIABLE)` -> hardcoded values

	RouterOS has predefined variables that can be used:
    - `HOSTNAME` - client hostname
    - `RADIUS_MT_STR1` - from radius MT attr nr. `24`
    - `RADIUS_MT_STR2` - from radius MT attr nr. `25`
    - `REMOTE_ID` - agent remote-id
    - `NETWORK_GATEWAY - the first gateway from `/ip dhcp-server network`, note that this option won't work if used from lease.

Now it is also possible to combine data types into one, for example: `0x01'vards'$(HOSTNAME)`For example if HOSTNAME is 'kvm', then raw value will be 0x0176617264736b766d.

### Read-Only

- `id` (String) The ID of this resource.
- `raw_value` (String) Read-only field which shows raw DHCP option value (the format actually sent out).

## Import
Import is supported using the following syntax:
```shell
#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ipv6/dhcp/server/option get [print show-ids]]
terraform import routeros_ipv6_dhcp_server_option.test *3
#Or you can import a resource using one of its attributes
terraform import routeros_ipv6_dhcp_server_option.test "name=domain-search"
```
