ddns-update-style none;
ignore client-updates;
ddns-update-style none;
ignore client-updates;

next-server ${BMP_HOST_IP};
filename \"pxelinux.0\";

default-lease-time         600;
max-lease-time             1200;
option domain-name \"localhost\";
option domain-name-servers 114.114.114.114,8.8.8.8;

#pxe
option client-system-architecture code 93 = unsigned integer 16;
class \"pxe-clients\" {
    match if (substring (option vendor-class-identifier, 0, 9) = \"PXEClient\") or
             (substring (option vendor-class-identifier, 0, 9) = \"HW-Client\");
    if option client-system-architecture = 00:00 {
        filename \"pxelinux.0\";
    }else if option client-system-architecture = 00:07 {
        filename \"/uefi/x86_64/grubx64.efi\";
    }else if option client-system-architecture = 00:0b {
        filename \"/uefi/arm64/grubaa64.efi\";
    }else if option client-system-architecture = 00:0c {
        filename \"/uefi/loongarch64/BOOTLOONGARCH64.EFI\";
    }else if option client-system-architecture = 00:27 {
        filename \"/uefi/loongarch64/BOOTLOONGARCH64.EFI\";
    }
}

key omapi_key {
	algorithm hmac-md5;
	secret ${BMP_OMAPI_KEY};
};
omapi-key omapi_key;
omapi-port ${BMP_OMAPI_PORT};

subnet ${BMP_HOST_IP} netmask 255.255.255.255 {
    
}

