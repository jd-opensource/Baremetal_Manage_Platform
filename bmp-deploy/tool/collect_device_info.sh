#!/bin/bash

RAID_DRIVER_ALL=(sas2ircu sas3ircu megacli64 storcli64 perccli64 arcconf no_raid)

SCRIPT_NAME=${0##*/}
usage(){    
    help="
Usage:  ./${SCRIPT_NAME} --sn SN --iloIp iloIp --iloUser iloUser --iloPassword iloPassword --ip ip --mask mask --gateway gateway --mac MAC [--raidDriver raidDriver]

"
    printf "$help"
}

parse_parameters()
{
    #parse parameters
    POSITIONAL_ARGS=()
    while [[ $# -gt 0 ]]; do
      case $1 in
        -h|--help)
          usage
          exit 0
          ;;
        --sn)
          sn="$2"
          shift
          shift
          ;;
        --iloIp)
          iloIp="$2"
          shift
          shift
          ;;
        --iloUser)
          iloUser="$2"
          shift
          shift
          ;;
        --iloPassword)
          iloPassword="$2"
          shift
          shift
          ;;
        --ip)
          ip="$2"
          shift
          shift
          ;;
        --mask)
          mask="$2"
          shift
          shift
          ;;
        --gateway)
          gateway="$2"
          shift
          shift
          ;;
        --mac)
          mac="$2"
          shift
          shift
          ;;
        --raidDriver)
          raidDriver="$2"
          shift
          shift
          ;;
        -*|--*)
          echo "Unknown option $1"
          exit 1
          ;;
        *)
          POSITIONAL_ARGS+=("$1") # save positional arg
          shift # past argument
          ;;
      esac
    done
    
    #check parameters
    [ -n "$sn" ] || { echo "sn is null"; usage; exit 1; }
    [ -n "$iloIp" ] || { echo "iloIp is null"; usage; exit 1; }
    [ -n "$iloUser" ] || { echo "iloUser is null"; usage; exit 1; }
    [ -n "$iloPassword" ] || { echo "iloPassword is null"; usage; exit 1; }
    [ -n "$ip" ] || { echo "ip is null"; usage; exit 1; }
    [ -n "$mask" ] || { echo "mask is null"; usage; exit 1; }
    [ -n "$gateway" ] || { echo "gateway is null"; usage; exit 1; }
    [ -n "$mac" ] || { echo "mac is null"; usage; exit 1; }    
#    [ -n "$raidDriver" ] || raidDriver="megacli64"
    [ -n "$raidDriver" ] && [[ ! " ${RAID_DRIVER_ALL[@]} " =~ " $raidDriver " ]] && { echo "raidDriver $raidDriver is not supported! must be one of ${RAID_DRIVER_ALL[@]}"; exit 1; }
}

parse_parameters $@
cat << EOF | curl -X POST -H "Content-Type: application/json; charset=utf-8" --data-binary @- 'http://127.0.0.1:8801/v1/collect/collectDeviceInfo'
{"collects":[{"sn":"$sn", "raidDriver":"$raidDriver", "allowOverride":true, "adapterId":0, "iloIp":"$iloIp", "iloUser":"$iloUser", "iloPassword":"$iloPassword", "mac1":"$mac", "privateIpv4":"$ip", "mask":"$mask","gateway":"$gateway"}]}
EOF

